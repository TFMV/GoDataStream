package main

import (
	"log"

	"github.com/TFMV/GoDataStream/consumer/storage"
)

func main() {
	consumer := storage.NewKafkaConsumer([]string{"localhost:9092"}, "user-topic", "example-group")
	defer consumer.Close()

	consumer.Consume(func(key, value []byte) {
		log.Printf("received message: key=%s value=%s", string(key), string(value))
		storage.StoreInBigQuery("your_dataset_id", "your_table_id", value)
	})
}
