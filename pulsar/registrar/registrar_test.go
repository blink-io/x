package registrar

import (
	"context"
	"testing"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/stretchr/testify/require"
)

func TestIface(t *testing.T) {
	ctx := context.Background()
	var rr ServiceRegistrar
	c, err := pulsar.NewClient(pulsar.ClientOptions{})
	require.NoError(t, err)
	rr = c
	require.NotNil(t, rr)

	str := "Hello"
	r := New[string](str, func(rr ServiceRegistrar, s string) {

	})

	err = r.RegisterToPulsar(ctx, c)
	require.NoError(t, err)
}
