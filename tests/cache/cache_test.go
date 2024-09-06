package cache

import (
	"testing"

	cache "github.com/Code-Hex/go-generics-cache"
	"github.com/stretchr/testify/require"
)

func TestCache_1(t *testing.T) {
	cc := cache.New[string, string]()
	cc.Set("kv", "good")

	vv, ok := cc.Get("kv")
	require.True(t, ok)
	require.Equal(t, "good", vv)
}
