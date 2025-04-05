package main

import (
	"api-gateway/handler"
	"api-gateway/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Настройка маршрутов
	r := gin.Default()

	// Применение middleware
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.AuthMiddleware)

	r.RedirectFixedPath = false

	// Роуты для API Gateway
	handler.RegisterRoutes(r)

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
