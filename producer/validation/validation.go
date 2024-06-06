package validation

import (
	"regexp"

	"github.com/TFMV/GoDataStream/producer/models"
)

// ValidateUser checks if the user record is valid
func ValidateUser(user models.User) bool {
	if user.Id == "" || user.Name == "" || user.Email == "" {
		return false
	}
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(user.Email)
}
