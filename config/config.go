package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Kafka struct {
		Brokers []string
		Topic   string
		GroupID string
	}
	BigQuery struct {
		ProjectID string
		DatasetID string
		TableID   string
	}
	Arrow struct {
		FilePath string
	}
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}
}
