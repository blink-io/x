package goredis

import (
	"context"
	"time"

	"github.com/blink-io/x/cache"
	"github.com/redis/go-redis/v9"
)

const Name = "goredis"

var _ cache.TTLCache[any] = (*Cache[any])(nil)

type Cache[V any] struct {
	rc  redis.UniversalClient
	ttl time.Duration
	ctx context.Context
}

func New[V any](rc redis.UniversalClient, ttl time.Duration) (*Cache[V], error) {
	return &Cache[V]{
		rc:  rc,
		ttl: ttl,
		ctx: context.Background(),
	}, nil
}

func (c *Cache[V]) Set(key string, value V) {
	c.rc.Set(c.ctx, key, value, c.ttl)
}

func (c *Cache[V]) SetWithTTL(key string, value V, ttl time.Duration) {
	c.rc.Set(c.ctx, key, value, ttl)
}

func (c *Cache[V]) Get(key string) (V, bool) {
	var v V

	cmd := c.rc.Get(c.ctx, key)
	if err := cmd.Err(); err != nil {
		if err == redis.Nil {
			return v, false
		}
		return v, false
	}

	if err := cmd.Scan(&v); err != nil {
		var zero V
		return zero, false
	}

	return v, true
}

func (c *Cache[V]) Del(key string) {
	c.rc.Del(c.ctx, key)
}

func (c *Cache[V]) Name() string {
	return Name
}
