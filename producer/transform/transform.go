package transform

import (
	"encoding/json"
	"strings"
)

// TransformData applies transformations to a data record
func TransformData(data []byte) []byte {
	var record map[string]interface{}
	json.Unmarshal(data, &record)

	if name, ok := record["name"].(string); ok {
		record["name"] = strings.ToUpper(name)
	}
	if email, ok := record["email"].(string); ok {
		record["email"] = strings.ToLower(email)
	}

	transformedData, _ := json.Marshal(record)
	return transformedData
}
