// Package rueidis contains the Redis store implementation.
package rueidis

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/blink-io/x/kvstore"
	. "github.com/blink-io/x/kvstore/store/redis/shared"

	"github.com/redis/rueidis"
)

// Name the name of the store.
const Name = "rueidis"

var _ kvstore.Store = (*Store)(nil)

// registers Redis to kvstore.
func init() {
	kvstore.Register(Name, newStore)
}

func newStore(ctx context.Context, endpoints []string, options kvstore.Config) (kvstore.Store, error) {
	cfg, ok := options.(*Config)
	if !ok && options != nil {
		return nil, &kvstore.InvalidConfigurationError{Store: Name, Config: options}
	}

	return New(ctx, endpoints, cfg)
}

// Store implements the store.Store interface.
type Store struct {
	client rueidis.Client
	script *rueidis.Lua
	codec  Codec
}

// New creates a new Redis client.
func New(ctx context.Context, endpoints []string, options *Config) (*Store, error) {
	return NewWithCodec(ctx, endpoints, options, &RawCodec{})
}

// NewWithCodec creates a new Redis client with codec config.
func NewWithCodec(ctx context.Context, endpoints []string, options *Config, codec Codec) (*Store, error) {
	client, err := newClient(endpoints, options)
	if err != nil {
		return nil, err
	}

	return makeStore(ctx, client, codec), nil
}

func newClient(endpoints []string, cfg *Config) (rueidis.Client, error) {
	if len(endpoints) > 1 {
		return nil, ErrMultipleEndpointsUnsupported
	}

	opt := rueidis.ClientOption{
		InitAddress: endpoints,
		//DialTimeout:  5 * time.Second,
		//ReadTimeout:  30 * time.Second,
		//WriteTimeout: 30 * time.Second,
	}

	if cfg != nil {
		opt.TLSConfig = cfg.TLS
		opt.Username = cfg.Username
		opt.Password = cfg.Password
		opt.SelectDB = cfg.DB
	}

	if cfg != nil && cfg.Sentinel != nil {
		if cfg.Sentinel.MasterName == "" {
			return nil, ErrMasterSetMustBeProvided
		}

		if !cfg.Sentinel.ClusterClient && (cfg.Sentinel.RouteByLatency || cfg.Sentinel.RouteRandomly) {
			return nil, ErrInvalidRoutesOptions
		}

		// TODO Need to be verified
		opt.Sentinel = rueidis.SentinelOption{
			Username:  cfg.Sentinel.Username,
			Password:  cfg.Sentinel.Password,
			MasterSet: cfg.Sentinel.MasterName,
			TLSConfig: opt.TLSConfig,
		}
	}

	return rueidis.NewClient(opt)
}

func makeStore(ctx context.Context, client rueidis.Client, codec Codec) *Store {
	// Listen to Keyspace events.
	cfgsetCmd := client.B().ConfigSet().
		ParameterValue().
		ParameterValue(ConfigSetParam, ConfigSetVal).
		Build()
	if err := client.Do(ctx, cfgsetCmd).Error(); err != nil {
		log.Printf("unable to set config value for: %s", ConfigSetParam)
	}

	var c Codec = &JSONCodec{}
	if codec != nil {
		c = codec
	}

	return &Store{
		client: client,
		script: rueidis.NewLuaScript(LuaScript()),
		codec:  c,
	}
}

// Put a value at the specified key.
func (r *Store) Put(ctx context.Context, key string, value []byte, opts *kvstore.WriteOptions) error {
	expirationAfter := NoExpiration
	if opts != nil && opts.TTL != 0 {
		expirationAfter = opts.TTL
	}

	return r.setTTL(ctx, normalize(key), &kvstore.KVPair{
		Key:       key,
		Value:     value,
		LastIndex: sequenceNum(),
	}, expirationAfter)
}

func (r *Store) setTTL(ctx context.Context, key string, val *kvstore.KVPair, ttl time.Duration) error {
	valStr, err := r.codec.Encode(val)
	if err != nil {
		return err
	}
	setCmd := r.client.B().Set().Key(key).Value(valStr).Ex(ttl).Build()
	return r.client.Do(ctx, setCmd).Error()
}

// Get a value given its key.
func (r *Store) Get(ctx context.Context, key string, _ *kvstore.ReadOptions) (*kvstore.KVPair, error) {
	return r.get(ctx, normalize(key))
}

func (r *Store) get(ctx context.Context, key string) (*kvstore.KVPair, error) {
	getCmd := r.client.B().Get().Key(key).Build()
	reply, err := r.client.Do(ctx, getCmd).AsBytes()
	if err != nil {
		if rueidis.IsRedisNil(err) {
			return nil, kvstore.ErrKeyNotFound
		}
		return nil, err
	}
	val := kvstore.KVPair{}
	if err := r.codec.Decode(reply, &val); err != nil {
		return nil, err
	}

	if val.Key == "" {
		val.Key = key
	}

	return &val, nil
}

// Delete the value at the specified key.
func (r *Store) Delete(ctx context.Context, key string) error {
	delCmd := r.client.B().Del().Key(normalize(key)).Build()
	return r.client.Do(ctx, delCmd).Error()
}

// Exists verify if a Key exists in the store.
func (r *Store) Exists(ctx context.Context, key string, _ *kvstore.ReadOptions) (bool, error) {
	existsCmd := r.client.B().Exists().Key(normalize(key)).Build()
	count, err := r.client.Do(ctx, existsCmd).AsInt64()
	return count != 0, err
}

// Watch for changes on a key.
// glitch: we use notified-then-retrieve to retrieve *kvstore.KVPair.
// so the responses may sometimes inaccurate.
func (r *Store) Watch(ctx context.Context, key string, _ *kvstore.ReadOptions) (<-chan *kvstore.KVPair, error) {
	watchCh := make(chan *kvstore.KVPair)
	nKey := normalize(key)

	get := getter(func() (interface{}, error) {
		pair, err := r.get(ctx, nKey)
		if err != nil {
			return nil, err
		}
		return pair, nil
	})

	push := pusher(func(v interface{}) {
		if val, ok := v.(*kvstore.KVPair); ok {
			watchCh <- val
		}
	})

	sub := newSubscribe(ctx, r.client, regexWatch(nKey, true))

	go func(ctx context.Context, sub *subscribe, get getter, push pusher) {
		defer func() {
			close(watchCh)
			_ = sub.Close()
		}()

		msgCh := sub.Receive(ctx)
		if err := watchLoop(ctx, msgCh, get, push); err != nil {
			log.Printf("watchLoop in Watch err: %v", err)
		}

	}(ctx, sub, get, push)

	return watchCh, nil
}

// WatchTree watches for changes on child nodes under a given directory.
func (r *Store) WatchTree(ctx context.Context, directory string, _ *kvstore.ReadOptions) (<-chan []*kvstore.KVPair, error) {
	watchCh := make(chan []*kvstore.KVPair)
	nKey := normalize(directory)

	get := getter(func() (interface{}, error) {
		pair, err := r.list(ctx, nKey)
		if err != nil {
			return nil, err
		}
		return pair, nil
	})

	push := pusher(func(v interface{}) {
		if p, ok := v.([]*kvstore.KVPair); ok {
			watchCh <- p
		}
	})

	sub := newSubscribe(ctx, r.client, regexWatch(nKey, true))

	go func(ctx context.Context, sub *subscribe, get getter, push pusher) {
		defer func() {
			close(watchCh)
			_ = sub.Close()
		}()

		msgCh := sub.Receive(ctx)
		if err := watchLoop(ctx, msgCh, get, push); err != nil {
			log.Printf("watchLoop in WatchTree err:%v\n", err)
		}
	}(ctx, sub, get, push)

	return watchCh, nil
}

// NewLock creates a lock for a given key.
// The returned Locker is not held and must be acquired
// with `.Lock`. The Value is optional.
func (r *Store) NewLock(_ context.Context, key string, opts *kvstore.LockOptions) (kvstore.Locker, error) {
	ttl := DefaultLockTTL
	var value []byte

	if opts != nil {
		if opts.TTL != 0 {
			ttl = opts.TTL
		}

		if len(opts.Value) != 0 {
			value = opts.Value
		}
	}

	return &redisLock{
		redis:    r,
		last:     nil,
		key:      key,
		value:    value,
		ttl:      ttl,
		unlockCh: make(chan struct{}),
	}, nil
}

// List the content of a given prefix.
func (r *Store) List(ctx context.Context, directory string, _ *kvstore.ReadOptions) ([]*kvstore.KVPair, error) {
	return r.list(ctx, normalize(directory))
}

func (r *Store) list(ctx context.Context, directory string) ([]*kvstore.KVPair, error) {
	regex := scanRegex(directory) // for all keyed with $directory.
	allKeys, err := r.keys(ctx, regex)
	if err != nil {
		return nil, err
	}

	// TODO: need to handle when #key is too large.
	return r.mget(ctx, directory, allKeys...)
}

func (r *Store) keys(ctx context.Context, regex string) ([]string, error) {
	const (
		startCursor  = 0
		endCursor    = 0
		defaultCount = 10
	)

	var allKeys []string

	doScanFn := func(ctx context.Context, regex string, cursor uint64, count int64) ([]string, uint64, error) {
		scanCmd := r.client.B().Scan().Cursor(cursor).Match(regex).Count(count).Build()
		scanEntry, doScanErr := r.client.Do(ctx, scanCmd).AsScanEntry()
		if doScanErr != nil {
			return nil, 0, doScanErr
		}
		return scanEntry.Elements, scanEntry.Cursor, nil
	}

	keys, nextCursor, err := doScanFn(ctx, regex, startCursor, defaultCount)
	if err != nil {
		return nil, err
	}

	allKeys = append(allKeys, keys...)

	for nextCursor != endCursor {
		keys, nextCursor, err = doScanFn(ctx, regex, nextCursor, defaultCount)
		if err != nil {
			return nil, err
		}

		allKeys = append(allKeys, keys...)
	}

	if len(allKeys) == 0 {
		return nil, kvstore.ErrKeyNotFound
	}

	return allKeys, nil
}

// mget values given their keys.
func (r *Store) mget(ctx context.Context, directory string, keys ...string) ([]*kvstore.KVPair, error) {
	mgetCmd := r.client.B().Mget().Key(keys...).Build()
	replies, err := r.client.Do(ctx, mgetCmd).AsStrSlice()
	if err != nil {
		return nil, err
	}

	var pairs []*kvstore.KVPair
	for i, reply := range replies {
		if reply == "" {
			// empty reply.
			continue
		}

		pair := &kvstore.KVPair{}
		if err := r.codec.Decode([]byte(reply), pair); err != nil {
			return nil, err
		}

		if pair.Key == "" {
			pair.Key = keys[i]
		}

		if normalize(pair.Key) != directory {
			pairs = append(pairs, pair)
		}
	}
	return pairs, nil
}

// DeleteTree deletes a range of keys under a given directory.
// glitch: we list all available keys first and then delete them all
// it costs two operations on redis, so is not atomicity.
func (r *Store) DeleteTree(ctx context.Context, directory string) error {
	regex := scanRegex(normalize(directory)) // for all keyed with $directory.

	allKeys, err := r.keys(ctx, regex)
	if err != nil {
		return err
	}

	delCmd := r.client.B().Del().Key(allKeys...).Build()
	return r.client.Do(ctx, delCmd).Error()
}

// AtomicPut is an atomic CAS operation on a single value.
// Pass previous = nil to create a new key.
// We introduced script on this page, so atomicity is guaranteed.
func (r *Store) AtomicPut(ctx context.Context, key string, value []byte, previous *kvstore.KVPair, opts *kvstore.WriteOptions) (bool, *kvstore.KVPair, error) {
	expirationAfter := NoExpiration
	if opts != nil && opts.TTL != 0 {
		expirationAfter = opts.TTL
	}

	newKV := &kvstore.KVPair{
		Key:       key,
		Value:     value,
		LastIndex: sequenceNum(),
	}
	nKey := normalize(key)

	// if previous == nil, set directly.
	if previous == nil {
		if err := r.setNX(ctx, nKey, newKV, expirationAfter); err != nil {
			return false, nil, err
		}
		return true, newKV, nil
	}

	if err := r.cas(ctx, nKey, previous, newKV, formatSec(expirationAfter)); err != nil {
		return false, nil, err
	}
	return true, newKV, nil
}

func (r *Store) setNX(ctx context.Context, key string, val *kvstore.KVPair, expirationAfter time.Duration) error {
	valStr, err := r.codec.Encode(val)
	if err != nil {
		return err
	}
	setCmdB := r.client.B().Set().Key(key).Value(valStr)
	if expirationAfter > 0 {
		setCmdB.Nx().Ex(expirationAfter)
	}
	ok, err := r.client.Do(ctx, setCmdB.Build()).AsBool()
	if err != nil {
		return err
	}
	if !ok {
		return kvstore.ErrKeyExists
	}
	return nil
}

func (r *Store) cas(ctx context.Context, key string, oldPair, newPair *kvstore.KVPair, secInStr string) error {
	newVal, err := r.codec.Encode(newPair)
	if err != nil {
		return err
	}

	oldVal, err := r.codec.Encode(oldPair)
	if err != nil {
		return err
	}

	return r.runScript(ctx, CmdCAS, key, oldVal, newVal, secInStr)
}

// AtomicDelete is an atomic delete operation on a single value
// the value will be deleted if previous matched the one stored in db.
func (r *Store) AtomicDelete(ctx context.Context, key string, previous *kvstore.KVPair) (bool, error) {
	if err := r.cad(ctx, normalize(key), previous); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Store) cad(ctx context.Context, key string, old *kvstore.KVPair) error {
	oldVal, err := r.codec.Encode(old)
	if err != nil {
		return err
	}

	return r.runScript(ctx, CmdCAD, key, oldVal)
}

// Close the store connection.
func (r *Store) Close() error {
	r.client.Close()
	return nil
}

func (r *Store) runScript(ctx context.Context, args ...string) error {
	err := r.script.Exec(ctx, r.client, nil, args).Error()
	if err != nil && strings.Contains(err.Error(), "redis: key is not found") {
		return kvstore.ErrKeyNotFound
	}
	if err != nil && strings.Contains(err.Error(), "redis: value has been changed") {
		return kvstore.ErrKeyModified
	}
	return err
}

func regexWatch(key string, withChildren bool) string {
	if withChildren {
		// For all database and keys with $key prefix.
		return fmt.Sprintf("__keyspace*:%s*", key)
	}
	// For all database and keys with $key.
	return fmt.Sprintf("__keyspace*:%s", key)
}

// getter defines a func type which retrieves data from remote storage.
type getter func() (interface{}, error)

// pusher defines a func type which pushes data blob into watch channel.
type pusher func(interface{})

func watchLoop(ctx context.Context, msgCh chan *rueidis.PubSubMessage, get getter, push pusher) error {
	// deliver the original data before we set up any events.
	pair, err := get()
	if err != nil && !errors.Is(err, kvstore.ErrKeyNotFound) {
		return err
	}

	if errors.Is(err, kvstore.ErrKeyNotFound) {
		pair = &kvstore.KVPair{}
	}

	push(pair)

	for m := range msgCh {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// retrieve and send back.
		pair, err := get()
		if err != nil && !errors.Is(err, kvstore.ErrKeyNotFound) {
			return err
		}

		// in case of watching a key that has been expired or deleted return and empty KV.
		if errors.Is(err, kvstore.ErrKeyNotFound) && (m.Message == "expired" || m.Message == "del") {
			pair = &kvstore.KVPair{}
		}

		push(pair)
	}

	return nil
}

type subscribe struct {
	client  rueidis.Client
	pubsub  rueidis.Completed
	closeCh chan struct{}
}

func newSubscribe(ctx context.Context, client rueidis.Client, regex string) *subscribe {
	return &subscribe{
		client:  client,
		pubsub:  client.B().Psubscribe().Pattern(regex).Build(),
		closeCh: make(chan struct{}),
	}
}

func (s *subscribe) Close() error {
	close(s.closeCh)
	s.client.Close()
	return nil
}

func (s *subscribe) Receive(ctx context.Context) chan *rueidis.PubSubMessage {
	msgCh := make(chan *rueidis.PubSubMessage)
	go s.receiveLoop(ctx, msgCh)
	return msgCh
}

func (s *subscribe) receiveLoop(ctx context.Context, msgCh chan *rueidis.PubSubMessage) {
	defer close(msgCh)

	for {
		select {
		case <-s.closeCh:
			return
		case <-ctx.Done():
			return
		default:
			var msg *rueidis.PubSubMessage
			err := s.client.Receive(ctx, s.pubsub, func(rmsg rueidis.PubSubMessage) {
				msg = &rmsg
			})
			if err != nil {
				return
			}
			if msg != nil {
				msgCh <- msg
			}
		}
	}
}

type redisLock struct {
	redis    *Store
	last     *kvstore.KVPair
	unlockCh chan struct{}

	key   string
	value []byte
	ttl   time.Duration
}

func (l *redisLock) Lock(ctx context.Context) (<-chan struct{}, error) {
	lockHeld := make(chan struct{})

	success, err := l.tryLock(ctx, lockHeld)
	if err != nil {
		return nil, err
	}
	if success {
		return lockHeld, nil
	}

	// wait for changes on the key.
	watch, err := l.redis.Watch(ctx, l.key, nil)
	if err != nil {
		return nil, err
	}

	for {
		select {
		case <-ctx.Done():
			return nil, ErrAbortTryLock
		case <-watch:
			success, err := l.tryLock(ctx, lockHeld)
			if err != nil {
				return nil, err
			}
			if success {
				return lockHeld, nil
			}
		}
	}
}

// tryLock return `true, nil` when it acquired and hold the lock
// and return `false, nil` when it can't lock now,
// and return `false, err` if any unexpected error happened underlying.
func (l *redisLock) tryLock(ctx context.Context, lockHeld chan struct{}) (bool, error) {
	success, item, err := l.redis.AtomicPut(ctx, l.key, l.value, l.last, &kvstore.WriteOptions{
		TTL: l.ttl,
	})
	if success {
		l.last = item
		// keep holding.
		go l.holdLock(ctx, lockHeld)
		return true, nil
	}
	if errors.Is(err, kvstore.ErrKeyNotFound) || errors.Is(err, kvstore.ErrKeyModified) || errors.Is(err, kvstore.ErrKeyExists) {
		return false, nil
	}
	return false, err
}

func (l *redisLock) holdLock(ctx context.Context, lockHeld chan struct{}) {
	defer close(lockHeld)

	hold := func() error {
		_, item, err := l.redis.AtomicPut(ctx, l.key, l.value, l.last, &kvstore.WriteOptions{
			TTL: l.ttl,
		})
		if err == nil {
			l.last = item
		}
		return err
	}

	heartbeat := time.NewTicker(l.ttl / 3)
	defer heartbeat.Stop()

	for {
		select {
		case <-heartbeat.C:
			if err := hold(); err != nil {
				return
			}
		case <-l.unlockCh:
			return
		case <-ctx.Done():
			return
		}
	}
}

func (l *redisLock) Unlock(ctx context.Context) error {
	l.unlockCh <- struct{}{}

	_, err := l.redis.AtomicDelete(ctx, l.key, l.last)
	if err != nil {
		return err
	}

	l.last = nil

	return nil
}

func scanRegex(directory string) string {
	return fmt.Sprintf("%s*", directory)
}

func normalize(key string) string {
	return strings.TrimPrefix(key, "/")
}

func formatSec(dur time.Duration) string {
	return fmt.Sprintf("%d", int(dur/time.Second))
}

func sequenceNum() uint64 {
	// TODO: use uuid if we concerns collision probability of this number.
	return uint64(time.Now().Nanosecond())
}
