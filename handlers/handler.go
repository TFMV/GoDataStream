package handlers

import (
	"net/http"

	"github.com/TFMV/GoDataStream/storage"
	"github.com/TFMV/GoDataStream/transform"
	"github.com/TFMV/GoDataStream/validation"

	"github.com/TFMV/GoDataStream/models"

	"github.com/gin-gonic/gin"
	flatbuffers "github.com/google/flatbuffers/go"
)

func EncodeUser(c *gin.Context) {
	var user map[string]interface{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validation.ValidateData(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user = transform.TransformData(user)

	builder := flatbuffers.NewBuilder(0)
	name := builder.CreateString(user["name"].(string))
	email := builder.CreateString(user["email"].(string))

	models.UserStart(builder)
	models.UserAddId(builder, int32(user["id"].(float64)))
	models.UserAddName(builder, name)
	models.UserAddEmail(builder, email)
	userOffset := models.UserEnd(builder)
	builder.Finish(userOffset)

	storage.StoreInBigQuery("your_dataset_id", "your_table_id", user)

	c.Data(http.StatusOK, "application/octet-stream", builder.FinishedBytes())
}

func DecodeUser(c *gin.Context) {
	buf := make([]byte, c.Request.ContentLength)
	c.Request.Body.Read(buf)
	user := models.GetRootAsUser(buf, 0)

	data := map[string]interface{}{
		"id":    user.Id(),
		"name":  string(user.Name()),
		"email": string(user.Email()),
	}

	c.JSON(http.StatusOK, data)
}
