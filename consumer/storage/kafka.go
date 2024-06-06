package storage

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

// KafkaConsumer represents a Kafka consumer
type KafkaConsumer struct {
	reader *kafka.Reader
}

// NewKafkaConsumer creates a new Kafka consumer
func NewKafkaConsumer(brokers []string, topic, groupID string) *KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		GroupID:     groupID,
		Topic:       topic,
		StartOffset: kafka.FirstOffset,
		MinBytes:    10e3, // 10KB
		MaxBytes:    10e6, // 10MB
	})
	return &KafkaConsumer{reader: reader}
}

// Consume reads messages from the Kafka topic
func (kc *KafkaConsumer) Consume(handler func(key, value []byte)) {
	for {
		msg, err := kc.reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("could not read message: %v", err)
			break
		}
		handler(msg.Key, msg.Value)
	}
}

// Close closes the Kafka consumer
func (kc *KafkaConsumer) Close() error {
	return kc.reader.Close()
}
