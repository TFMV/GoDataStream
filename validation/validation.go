package validation

import "errors"

func ValidateData(data map[string]interface{}) error {
	if data["id"] == nil || data["name"] == nil || data["email"] == nil {
		return errors.New("missing required fields")
	}
	return nil
}
