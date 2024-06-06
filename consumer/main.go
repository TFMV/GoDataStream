package main

import (
	"log"

	"github.com/TFMV/GoDataStream/config"
	kafka "github.com/TFMV/GoDataStream/consumer/kafka"
	storage "github.com/TFMV/GoDataStream/consumer/storage"
	"github.com/TFMV/GoDataStream/producer/transform"
)

func main() {
	config.LoadConfig()

	consumer := kafka.NewKafkaConsumer(config.AppConfig.Kafka.Brokers, config.AppConfig.Kafka.Topic, config.AppConfig.Kafka.GroupID)
	defer consumer.Close()

	consumer.Consume(func(key, value []byte) {
		log.Printf("received message: key=%s value=%s", string(key), string(value))

		transformedData := transform.TransformData(value)

		err := storage.StoreInBigQuery(config.AppConfig.BigQuery.ProjectID, config.AppConfig.BigQuery.DatasetID, config.AppConfig.BigQuery.TableID, transformedData)
		if err != nil {
			log.Printf("Error storing data in BigQuery: %v", err)
		}
	})
}
