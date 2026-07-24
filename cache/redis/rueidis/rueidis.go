package rueidis

import (
	"context"
	"errors"
	"time"

	"github.com/blink-io/x/cache"
	"github.com/blink-io/x/encoding"

	"github.com/redis/rueidis"
)

var errNilCodec = errors.New("rueidis cache: encoding codec is required")

const Name = "goredis"

var _ cache.TTLCache[any] = (*Cache[any])(nil)

type Cache[V any] struct {
	cc  rueidis.Client
	ttl time.Duration
	ctx context.Context
	enc encoding.Codec
}

// New returns a new rueidis-backed TTL cache. The provided encoding.Codec is
// used to serialize values into bytes for storage and to decode them on
// retrieval; it must be safe for concurrent use.
func New[V any](cc rueidis.Client, ttl time.Duration, enc encoding.Codec) (*Cache[V], error) {
	if enc == nil {
		return nil, errNilCodec
	}
	return &Cache[V]{
		cc:  cc,
		ttl: ttl,
		ctx: context.Background(),
		enc: enc,
	}, nil
}

func (c *Cache[V]) Set(key string, value V) {
	c.setWithTTL(key, value, c.ttl)
}

func (c *Cache[V]) SetWithTTL(key string, value V, ttl time.Duration) {
	c.setWithTTL(key, value, ttl)
}

func (c *Cache[V]) setWithTTL(key string, value V, ttl time.Duration) {
	data, err := c.enc.Marshal(value)
	if err != nil {
		return
	}
	cmd := c.cc.B().Set().Key(key).Value(string(data)).Ex(ttl).Build()
	c.cc.Do(c.ctx, cmd)
}

func (c *Cache[V]) Get(key string) (V, bool) {
	cmd := c.cc.B().Get().Key(key).Build()
	resp, err := c.cc.Do(c.ctx, cmd).AsBytes()
	if err != nil {
		var zero V
		return zero, false
	}
	var v V
	if err := c.enc.Unmarshal(resp, &v); err != nil {
		return v, false
	}
	return v, true
}

func (c *Cache[V]) Del(key string) {
	cmd := c.cc.B().Del().Key(key).Build()
	c.cc.Do(c.ctx, cmd)
}

func (c *Cache[V]) Name() string {
	return Name
}