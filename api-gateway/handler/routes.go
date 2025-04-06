package handler

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.Any("/products", proxyTo("http://localhost:8081", ""))
	r.Any("/products/*proxyPath", proxyTo("http://localhost:8081", ""))

	r.Any("/orders", proxyTo("http://localhost:8082", ""))
	r.Any("/orders/*proxyPath", proxyTo("http://localhost:8082", ""))
}

func proxyTo(targetHost string, _ string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("🚀 Proxy triggered! Method:", c.Request.Method, "Path:", c.Request.URL.Path)

		remote, err := url.Parse(targetHost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot parse backend URL"})
			return
		}

		c.Request.URL.Scheme = remote.Scheme
		c.Request.URL.Host = remote.Host
		c.Request.Host = remote.Host

		fmt.Println("➡️ Forwarding to:", remote.Host+c.Request.URL.Path)

		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
