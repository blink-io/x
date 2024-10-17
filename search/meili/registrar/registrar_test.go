package registrar

import (
	"context"
	"testing"

	"github.com/meilisearch/meilisearch-go"
	"github.com/stretchr/testify/require"
)

func TestIface(t *testing.T) {
	ctx := context.Background()
	var rr ServiceRegistrar
	s := meilisearch.New("", meilisearch.WithAPIKey(""))

	rr = NewServiceRegistrar(s)
	require.NotNil(t, rr)

	str := "Hello"
	r := New[string](str, func(rr ServiceRegistrar, s string) {

	})

	err := r.RegisterToMeiliSearch(ctx, rr)
	require.NoError(t, err)
}
