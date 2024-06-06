package main

import (
	"github.com/TFMV/GoDataStream/producer/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/encode", handlers.EncodeUser)
	r.Run(":8080") // Run on port 8080
}
