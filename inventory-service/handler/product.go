package handler

import (
	"inventory-service/model"
	"inventory-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Регистрация маршрутов для работы с продуктами
func RegisterRoutes(r *gin.Engine, svc *service.InventoryService) {
	r.POST("/products", createProduct(svc))
	r.GET("/products/:id", getProduct(svc))
	r.PATCH("/products/:id", updateProduct(svc))
	r.DELETE("/products/:id", deleteProduct(svc))
	r.GET("/products", listProducts(svc))
}

// Создание нового продукта
func createProduct(svc *service.InventoryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p model.Product
		if err := c.BindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		svc.Create(p)
		c.JSON(http.StatusCreated, p)
	}
}

// Получение продукта по ID
func getProduct(svc *service.InventoryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		product, ok := svc.Get(id)
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

// Обновление данных о продукте
func updateProduct(svc *service.InventoryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var p model.Product
		if err := c.BindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		svc.Update(id, p)
		c.JSON(http.StatusOK, p)
	}
}

// Удаление продукта
func deleteProduct(svc *service.InventoryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		svc.Delete(id)
		c.Status(http.StatusNoContent)
	}
}

// Список всех продуктов
func listProducts(svc *service.InventoryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		products := svc.List()
		c.JSON(http.StatusOK, products)
	}
}
