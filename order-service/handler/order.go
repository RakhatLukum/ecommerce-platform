package handler

import (
	"net/http"
	"order-service/model"
	"order-service/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, svc *service.OrderService) {
	r.POST("/orders", createOrder(svc))
	r.GET("/orders/:id", getOrder(svc))
	r.PATCH("/orders/:id", updateOrder(svc))
	r.GET("/orders", listOrders(svc))
}

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

func listOrders(svc *service.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		orders := svc.List()
		c.JSON(http.StatusOK, orders)
	}
}
