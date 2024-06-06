package storage

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

// KafkaProducer represents a Kafka producer
type KafkaProducer struct {
	writer *kafka.Writer
}

// NewKafkaProducer creates a new Kafka producer
func NewKafkaProducer(brokers []string, topic string) *KafkaProducer {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
	}
	return &KafkaProducer{writer: writer}
}

// Produce sends a message to the Kafka topic
func (kp *KafkaProducer) Produce(key, value []byte) error {
	msg := kafka.Message{
		Key:   key,
		Value: value,
	}
	return kp.writer.WriteMessages(context.Background(), msg)
}

// Close closes the Kafka producer
func (kp *KafkaProducer) Close() error {
	return kp.writer.Close()
}

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

// Example usage
func main() {
	brokers := []string{"localhost:9092"}
	topic := "example-topic"
	groupID := "example-group"

	// Create producer
	producer := NewKafkaProducer(brokers, topic)
	defer producer.Close()

	// Produce a message
	err := producer.Produce([]byte("key"), []byte("value"))
	if err != nil {
		log.Fatalf("could not produce message: %v", err)
	}

	// Create consumer
	consumer := NewKafkaConsumer(brokers, topic, groupID)
	defer consumer.Close()

	// Consume messages
	consumer.Consume(func(key, value []byte) {
		log.Printf("received message: key=%s value=%s", string(key), string(value))
	})
}
