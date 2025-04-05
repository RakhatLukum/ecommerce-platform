package main

import (
	"order-service/handler"
	"order-service/repository"
	"order-service/service"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := repository.NewMemoryRepo()
	orderService := service.NewOrderService(repo)

	r := gin.Default()

	handler.RegisterRoutes(r, orderService)

	r.Run(":8082")
}
