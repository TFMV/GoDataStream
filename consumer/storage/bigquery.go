package storage

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/bigquery"
)

// User represents a user record
type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Save implements the ValueSaver interface.
func (u *User) Save() (map[string]bigquery.Value, string, error) {
	return map[string]bigquery.Value{
		"id":    u.Id,
		"name":  u.Name,
		"email": u.Email,
	}, bigquery.NoDedupeID, nil
}

// StoreInBigQuery stores data in BigQuery using the legacy streaming insert API
func StoreInBigQuery(projectID, datasetID, tableID string, data []byte) error {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %w", err)
	}
	defer client.Close()

	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}

	inserter := client.Dataset(datasetID).Table(tableID).Inserter()
	if err := inserter.Put(ctx, []*User{&user}); err != nil {
		return fmt.Errorf("inserter.Put: %w", err)
	}

	fmt.Printf("Data inserted into BigQuery: %v\n", user)
	return nil
}
