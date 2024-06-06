package main

import (
	"log"

	"github.com/TFMV/GoDataStream/producer/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/encode", handlers.EncodeUser)

	log.Fatal(r.Run(":8080"))
}
