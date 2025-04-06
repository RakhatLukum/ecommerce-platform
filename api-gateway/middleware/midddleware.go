package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware(c *gin.Context) {
	start := time.Now()
	c.Next()
	duration := time.Since(start)
	log.Printf("Request %s %s took %v", c.Request.Method, c.Request.URL, duration)
}

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || authHeader != "Bearer token123" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	c.Next()
}
