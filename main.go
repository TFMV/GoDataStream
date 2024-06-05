package main

import (
	"fmt"

	"github.com/TFMV/GoDataStream/handlers"
	"github.com/TFMV/GoDataStream/models"
	"github.com/TFMV/GoDataStream/storage"
	"github.com/TFMV/GoDataStream/transform"
	"github.com/TFMV/GoDataStream/validation"
)

func main() {
	fmt.Println("GoDataStream microservice")
	handlers.HandleRequest()
	storage.StoreData()
	transform.TransformData()
	validation.ValidateData()
	models.PrintModel()
}
