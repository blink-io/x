package registrar

import (
	"context"
	"testing"

	"github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/require"
)

func TestIface(t *testing.T) {
	ctx := context.Background()
	var rr ServiceRegistrar
	c, err := amqp091.Dial("")
	require.NoError(t, err)

	require.NoError(t, err)

	rr = NewServiceRegistrar(c)
	require.NotNil(t, rr)

	str := "Hello"
	r := New[string](str, func(rr ServiceRegistrar, s string) {

	})

	err = r.RegisterToAMQPv091(ctx, rr)
	require.NoError(t, err)
}
