package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/TFMV/GoDataStream/producer/models"
	"github.com/TFMV/GoDataStream/producer/storage"
	"github.com/TFMV/GoDataStream/producer/transform"
	"github.com/TFMV/GoDataStream/producer/validation"
	"github.com/gin-gonic/gin"
)

func EncodeUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !validation.ValidateUser(user) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	transformedUser := transform.TransformUser(user)

	userBytes, err := json.Marshal(transformedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal user data"})
		return
	}

	storage.SendToKafka("user-topic", "user-key", userBytes)

	c.JSON(http.StatusOK, gin.H{"status": "User data sent to Kafka"})
}
