package rueidis

import (
	"context"
	"time"

	"github.com/blink-io/x/session/store"

	"github.com/redis/rueidis"
)

const Name = "rueidis"

var _ store.Store = (*Store)(nil)

// Store represents the session store.
type Store struct {
	client rueidis.Client
	prefix string
}

// New returns a new store instance. The client parameter should be a pointer
// to a go-redis connection.
func New(client rueidis.Client) *Store {
	return NewWithPrefix(client, store.DefaultPrefix)
}

// NewWithPrefix returns a new store instance. The pool parameter should be a pointer
// to a redigo connection pool. The prefix parameter controls the Redis key
// prefix, which can be used to avoid naming clashes if necessary.
func NewWithPrefix(client rueidis.Client, prefix string) *Store {
	return &Store{
		client: client,
		prefix: prefix,
	}
}

func (s *Store) Name() string {
	return Name
}

// Find returns the data for a given session token from the store instance.
// If the session token is not found or is expired, the returned exists flag
// will be set to false.
func (s *Store) Find(ctx context.Context, token string) (b []byte, exists bool, err error) {
	getCmd := s.client.B().Get().Key(s.prefix + token).Build()
	b, err = s.client.Do(ctx, getCmd).AsBytes()
	if rueidis.IsRedisNil(err) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}
	return b, true, nil
}

// Commit adds a session token and data to the store instance with the
// given expiry time. If the session token already exists then the data and
// expiry time are updated.
func (s *Store) Commit(ctx context.Context, token string, b []byte, expiry time.Time) error {
	setCmd := s.client.B().Set().Key(s.prefix + token).Value(string(b))
	if ttl := expiry.Sub(time.Now()); ttl > 0 {
		setCmd.Ex(ttl)
	}
	err := s.client.Do(ctx, setCmd.Build()).Error()
	return err
}

// Delete removes a session token and corresponding data from the store
// instance.
func (s *Store) Delete(ctx context.Context, token string) error {
	delCmd := s.client.B().Del().Key(s.prefix + token).Build()
	return s.client.Do(ctx, delCmd).Error()
}

// All returns a map containing the token and data for all active (i.e.
// not expired) sessions in the store instance.
func (s *Store) All(ctx context.Context) (map[string][]byte, error) {
	var cursor uint64
	sessions := make(map[string][]byte)

	for {
		// Batch one SCAN page of keys in a single MGET instead of one
		// sequential GET round-trip per session.
		scanCmd := s.client.B().Scan().Cursor(cursor).Match(s.prefix + "*").Count(500).Build()
		v, err := s.client.Do(ctx, scanCmd).AsScanEntry()
		if err != nil {
			if rueidis.IsRedisNil(err) {
				return nil, nil
			}
			return nil, err
		}
		cursor = v.Cursor
		keys := v.Elements

		if len(keys) > 0 {
			mgetCmd := s.client.B().Mget().Key(keys...).Build()
			replies, err := s.client.Do(ctx, mgetCmd).AsStrSlice()
			if err != nil {
				return nil, err
			}
			for i, reply := range replies {
				if reply != "" {
					sessions[keys[i][len(s.prefix):]] = []byte(reply)
				}
			}
		}
		if cursor == 0 {
			break
		}
	}
	return sessions, nil
}
