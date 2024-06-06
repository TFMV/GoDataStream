package transform

import (
	"strings"

	"github.com/TFMV/GoDataStream/models"
	flatbuffers "github.com/google/flatbuffers/go"
)

// TransformUser applies transformations to a user record
func TransformUser(user *models.User) *models.User {
	transformedName := strings.ToUpper(string(user.Name()))
	transformedEmail := strings.ToLower(string(user.Email()))

	builder := flatbuffers.NewBuilder(0)
	idOffset := builder.CreateByteVector(user.Id())
	nameOffset := builder.CreateString(transformedName)
	emailOffset := builder.CreateString(transformedEmail)

	models.UserStart(builder)
	models.UserAddId(builder, idOffset)
	models.UserAddName(builder, nameOffset)
	models.UserAddEmail(builder, emailOffset)
	userOffset := models.UserEnd(builder)

	builder.Finish(userOffset)
	return models.GetRootAsUser(builder.FinishedBytes(), 0)
}
