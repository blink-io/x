package pulsar

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/apache/pulsar-client-go/pulsar"
)

func TestExample_1(t *testing.T) {
	// Create a Pulsar client
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
	})
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	// Create a producer
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "topic-1",
	})
	if err != nil {
		log.Fatal(err)
	}

	defer producer.Close()

	ctx := context.Background()

	msgId, err := producer.Send(ctx, &pulsar.ProducerMessage{
		Payload: []byte(fmt.Sprintf("hello world")),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("msgId: ", msgId)
}
