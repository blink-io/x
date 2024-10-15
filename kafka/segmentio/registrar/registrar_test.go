package registrar

import (
	"context"
	"testing"

	kafkago "github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/require"
)

func TestIface(t *testing.T) {
	ctx := context.Background()
	var rr ServiceRegistrar
	c := &kafkago.Client{}
	rr = c
	require.NotNil(t, rr)

	str := "Hello"
	r := New[string](str, func(rr ServiceRegistrar, s string) {

	})

	err := r.RegisterToGoKafka(ctx, rr)
	require.NoError(t, err)
}
