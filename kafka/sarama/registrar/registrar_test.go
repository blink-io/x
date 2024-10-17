package registrar

import (
	"testing"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/require"
)

func TestIface(t *testing.T) {
	var rr ServiceRegistrar
	c, err := sarama.NewClient([]string{""}, &sarama.Config{})
	require.NoError(t, err)

	rr = NewServiceRegistrar(c)
	require.NotNil(t, rr)

	str := "Hello"
	r := New[string](str, func(rr ServiceRegistrar, s string) {

	})

	err = r.RegisterToSaramaKafka(nil, nil)
	require.NoError(t, err)
}
