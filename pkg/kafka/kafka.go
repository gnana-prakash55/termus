package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

// topic and brokers
const (
	topic   = "termus"
	brokers = "localhost:9092"
)

// @message the message from the server
// @ctx context provided
func Producer(ctx context.Context, message string) {

	// kafka writer
	w := kafka.NewWriter(kafka.WriterConfig{
		Topic:   topic,
		Brokers: []string{brokers},
	})

	// write message into kafka
	err := w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(message),
		Value: []byte(message),
	})

	// check error
	if err != nil {
		panic("could not write message " + err.Error())
	}

}
