package registrar

import (
	"context"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/stretchr/testify/require"
)

func TestIface(t *testing.T) {
	ctx := context.Background()
	var rr ServiceRegistrar
	c, err := elasticsearch.NewClient(elasticsearch.Config{})
	require.NoError(t, err)
	rr = NewServiceRegistrar(c)
	require.NotNil(t, rr)

	str := "Hello"
	r := New[string](str, func(rr ServiceRegistrar, s string) {

	})

	err = r.RegisterToElasticSearch(ctx, rr)
	require.NoError(t, err)
}
