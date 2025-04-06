package main

import (
	"api-gateway/handler"
	"api-gateway/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.AuthMiddleware)

	r.RedirectFixedPath = false

	handler.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
