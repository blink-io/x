package registrar

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/twmb/franz-go/pkg/kgo"
)

func TestIface(t *testing.T) {
	seeds := []string{"localhost:9092"}

	var rr ServiceRegistrar
	c, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
		kgo.ConsumerGroup("my-group-identifier"),
		kgo.ConsumeTopics("foo"),
	)
	require.NoError(t, err)

	rr = NewServiceRegistrar(c)
	require.NotNil(t, rr)

	str := "Hello"
	r := New[string](str, func(rr ServiceRegistrar, s string) {

	})

	err = r.RegisterToSaramaKafka(nil, nil)
	require.NoError(t, err)
}
