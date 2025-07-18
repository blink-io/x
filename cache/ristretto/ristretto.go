package ristretto

import (
	"time"

	"github.com/blink-io/x/cache"
	"github.com/dgraph-io/ristretto/v2"
)

const Name = "ristretto"
const defaultCost = 1

var _ cache.TTLCache[any] = (*Cache[any])(nil)

func init() {
	//local.SetProviderFn(ProviderLRU, NewLRULocal)
}

type Cache[V any] struct {
	cc  *ristretto.Cache[string, V]
	ttl time.Duration
}

func New[V any](ttl time.Duration) (*Cache[V], error) {
	c, err := ristretto.NewCache[string, V](&ristretto.Config[string, V]{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		return nil, err
	}
	return &Cache[V]{c, ttl}, nil
}

func (c *Cache[V]) Set(key string, value V) {
	c.cc.SetWithTTL(key, value, defaultCost, c.ttl)
}

func (c *Cache[V]) SetWithTTL(key string, value V, ttl time.Duration) {
	c.cc.SetWithTTL(key, value, defaultCost, ttl)
}

func (c *Cache[V]) Get(key string) (V, bool) {
	return c.cc.Get(key)
}

func (c *Cache[V]) Del(key string) {
	c.cc.Del(key)
}

func (c *Cache[V]) Name() string {
	return Name
}
