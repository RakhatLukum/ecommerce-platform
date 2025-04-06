package main

import (
	"inventory-service/handler"
	"inventory-service/repository"
	"inventory-service/service"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := repository.NewMemoryRepo()
	inventoryService := service.NewInventoryService(repo)

	r := gin.Default()

	handler.RegisterRoutes(r, inventoryService)

	r.Run(":8081")
}
