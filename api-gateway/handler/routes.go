package handler

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Исправленный прокси-запрос
	r.Any("/products/*proxyPath", proxyTo("http://localhost:8081", "/products"))
	r.Any("/orders/*proxyPath", proxyTo("http://localhost:8082", "/orders"))
}

func proxyTo(targetHost, basePath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Убираем префикс пути
		remote, err := url.Parse(targetHost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot parse backend URL"})
			return
		}

		// Исправление пути
		originalPath := c.Request.URL.Path
		trimmedPath := strings.TrimPrefix(originalPath, basePath)
		if trimmedPath == "" {
			trimmedPath = "/"
		}
		c.Request.URL.Path = trimmedPath
		c.Request.Host = remote.Host

		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
