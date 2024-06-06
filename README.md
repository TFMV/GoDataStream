# GoDataStream

GoDataStream is a high-performance Go microservice designed to showcase the speed and efficiency of Go for real-time data engineering tasks. This project demonstrates real-time data transformation, validation, and storage integration with BigQuery and Kafka.

![GDS](assets/GDS.webp)

## Features

- **Real-Time Data Transformation:** Efficiently transform user data using Go's powerful standard library.
- **Data Validation:** Validate user data to ensure integrity and consistency.
- **Kafka Integration:** Seamlessly send transformed user data to a Kafka topic.
- **BigQuery Integration:** Store validated and transformed user data into BigQuery for further analysis.

## Getting Started

### Prerequisites

- Go 1.16 or higher
- Kafka instance
- BigQuery setup

## Usage

### BigQuery Integration

Configure your BigQuery credentials and dataset/table information in consumer/storage/bigquery.go.

### Kafka Integration

Configure your Kafka brokers and topics in producer/storage/kafka.go and consumer/storage/kafka.go.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Author

Thomas F McGeehan V
