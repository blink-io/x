package registrar

import (
	"context"
	"testing"

	"github.com/Azure/go-amqp"
	"github.com/stretchr/testify/require"
)

func TestIface(t *testing.T) {
	ctx := context.Background()
	var rr ServiceRegistrar
	c, err := amqp.Dial(ctx, "amqp://localhost", &amqp.ConnOptions{})
	require.NoError(t, err)

	sess, err := c.NewSession(ctx, &amqp.SessionOptions{})
	require.NoError(t, err)

	rr = NewServiceRegistrar(sess)
	require.NotNil(t, rr)

	str := "Hello"
	r := New[string](str, func(rr ServiceRegistrar, s string) {

	})

	err = r.RegisterToAMQPv1(ctx, sess)
	require.NoError(t, err)
}
