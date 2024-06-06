package transform

import (
	"strings"

	"github.com/TFMV/GoDataStream/producer/models"
)

// TransformUser applies transformations to a user record
func TransformUser(user models.User) models.User {
	user.Name = strings.ToUpper(user.Name)
	user.Email = strings.ToLower(user.Email)
	return user
}
