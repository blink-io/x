package registrar

import (
	"context"
	"testing"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/stretchr/testify/require"
)

func TestIface(t *testing.T) {
	ctx := context.Background()
	c, err := opensearchapi.NewClient(opensearchapi.Config{})
	require.NoError(t, err)

	var rr ServiceRegistrar = c
	require.NotNil(t, rr)

	str := "Hello"
	r := New[string](str, func(rr ServiceRegistrar, s string) {

	})

	err = r.RegisterToOpenSearch(ctx, rr)
	require.NoError(t, err)
}
