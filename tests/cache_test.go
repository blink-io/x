package tests

import (
	"testing"

	"github.com/samber/hot"
)

func TestCache_Hot_1(t *testing.T) {
	cache := hot.NewHotCache[string, int](hot.LRU, 100_000).
		Build()

	cache.Set("hello", 42)
	cache.SetMany(map[string]int{"foo": 1, "bar": 2})
}
