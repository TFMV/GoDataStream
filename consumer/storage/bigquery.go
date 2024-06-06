package storage

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/bigquery"
)

// User represents a user record
type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// StoreInBigQuery stores data in BigQuery
func StoreInBigQuery(datasetID, tableID string, data []byte) error {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, "your-project-id")
	if err != nil {
		return err
	}
	defer client.Close()

	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		return err
	}

	u := client.Dataset(datasetID).Table(tableID).Uploader()
	if err := u.Put(ctx, user); err != nil {
		return err
	}

	log.Printf("Data inserted into BigQuery: %v", user)
	return nil
}
