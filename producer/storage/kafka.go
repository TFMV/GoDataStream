package storage

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

// SendToKafka sends data to the Kafka topic
func SendToKafka(topic, key string, value []byte) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: value,
		},
	)

	if err != nil {
		log.Fatalf("Failed to write message: %v", err)
		return err
	}

	w.Close()
	return nil
}
