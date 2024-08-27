package kafka

import (
	"testing"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/require"
)

func TestClient_1(t *testing.T) {
	addrs := []string{"127.0.0.1:9092"}

	cfg := sarama.NewConfig()
	cc, err := sarama.NewClient(addrs, cfg)
	require.NoError(t, err)
	require.NotNil(t, cc)

}
