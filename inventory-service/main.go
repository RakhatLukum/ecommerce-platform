package main

import (
	"inventory-service/handler"
	"inventory-service/repository"
	"inventory-service/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Создание репозитория и сервиса
	repo := repository.NewMemoryRepo()
	inventoryService := service.NewInventoryService(repo)

	// Создание Gin маршрутизатора
	r := gin.Default()

	// Роуты для Inventory Service
	handler.RegisterRoutes(r, inventoryService)

	// Запуск сервера
	r.Run(":8081")
}
