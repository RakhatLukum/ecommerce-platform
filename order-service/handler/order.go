package handler

import (
	"net/http"
	"order-service/model"
	"order-service/service"

	"github.com/gin-gonic/gin"
)

// Регистрация маршрутов для заказов
func RegisterRoutes(r *gin.Engine, svc *service.OrderService) {
	r.POST("/orders", createOrder(svc))
	r.GET("/orders/:id", getOrder(svc))
	r.PATCH("/orders/:id", updateOrder(svc))
	r.GET("/orders", listOrders(svc))
}

// Создание нового заказа
func createOrder(svc *service.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var o model.Order
		if err := c.BindJSON(&o); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		svc.Create(o)
		c.JSON(http.StatusCreated, o)
	}
}

// Получение заказа по ID
func getOrder(svc *service.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		order, ok := svc.Get(id)
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusOK, order)
	}
}

// Обновление статуса заказа
func updateOrder(svc *service.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var o model.Order
		if err := c.BindJSON(&o); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		svc.Update(id, o)
		c.JSON(http.StatusOK, o)
	}
}

// Список всех заказов
func listOrders(svc *service.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		orders := svc.List()
		c.JSON(http.StatusOK, orders)
	}
}
