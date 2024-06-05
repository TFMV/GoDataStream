package storage

import (
	"context"
	"log"

	"cloud.google.com/go/bigquery"
)

func StoreInBigQuery(datasetID, tableID string, data map[string]interface{}) error {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, "your-project-id")
	if err != nil {
		return err
	}
	defer client.Close()

	u := client.Dataset(datasetID).Table(tableID).Uploader()
	if err := u.Put(ctx, data); err != nil {
		return err
	}

	log.Printf("Data inserted into BigQuery: %v", data)
	return nil
}
