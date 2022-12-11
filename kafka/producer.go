package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func ProduceMessage(topic string, msg string) {
	w := &kafka.Writer{
		Addr:     kafka.TCP("kafka:9092"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("msg"),
			Value: []byte(msg),
		},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
