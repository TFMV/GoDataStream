package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/apache/arrow/go/arrow/ipc"
	"github.com/apache/arrow/go/arrow/memory"
	"github.com/segmentio/kafka-go"
)

func main() {
	// Open Arrow file
	file, err := os.Open("path/to/your/data.arrow")
	if err != nil {
		log.Fatalf("failed to open arrow file: %v", err)
	}
	defer file.Close()

	// Initialize Arrow reader
	mem := memory.NewGoAllocator()
	reader, err := ipc.NewFileReader(file, ipc.WithAllocator(mem))
	if err != nil {
		log.Fatalf("failed to create arrow reader: %v", err)
	}
	defer reader.Close()

	// Create a new Kafka producer
	brokers := []string{"localhost:9092"}
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    "user-topic",
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	// Read Arrow records and send to Kafka
	for i := 0; i < reader.NumRecords(); i++ {
		record, err := reader.Record(i)
		if err != nil {
			log.Fatalf("failed to read record: %v", err)
		}

		for j := 0; j < int(record.NumRows()); j++ {
			row := record.Column(j)
			value, err := json.Marshal(row)
			if err != nil {
				log.Printf("failed to marshal row: %v", err)
				continue
			}
			err = writer.WriteMessages(context.Background(),
				kafka.Message{
					Key:   []byte(fmt.Sprintf("%v", row)), // Adjust the key as per your schema
					Value: value,
				},
			)
			if err != nil {
				log.Printf("failed to write message to kafka: %v", err)
			}
		}
	}

	fmt.Println("Arrow data sent to Kafka")
}
