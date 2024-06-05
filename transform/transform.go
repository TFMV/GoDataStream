package transform

import "time"

func TransformData(data map[string]interface{}) map[string]interface{} {
	data["processed_at"] = time.Now().UTC().String()
	return data
}
