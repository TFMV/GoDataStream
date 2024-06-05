package main

import (
	"github.com/TFMV/GoDataStream/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/encode", handlers.EncodeUser)
	router.POST("/decode", handlers.DecodeUser)
	router.Run(":8080")
}
