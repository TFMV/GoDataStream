package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	fmt.Println("Go Data Stream")

	user := map[string]interface{}{
		"id":    1,
		"name":  "John Doe",
		"email": "",
	}

	// Convert user to []byte
	userBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshaling user:", err)
		return
	}

	// Create a new Kafka producer
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "user-topic",
		Balancer: &kafka.LeastBytes{},
	})

	w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("user"),
			Value: userBytes,
		},
	)

	w.Close()
	fmt.Println("User data sent to Kafka")

}
