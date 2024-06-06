package storage

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(brokers []string, topic string) *KafkaProducer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	return &KafkaProducer{writer: writer}
}

func (kp *KafkaProducer) Produce(key, value []byte) error {
	msg := kafka.Message{
		Key:   key,
		Value: value,
	}
	return kp.writer.WriteMessages(context.Background(), msg)
}

func (kp *KafkaProducer) Close() error {
	return kp.writer.Close()
}
