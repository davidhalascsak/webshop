package router

import (
	"net/http"
	"webshop/internal/product"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	ProductHandler *product.Handler
}

func Setup(h Handlers) *gin.Engine {
	engine := gin.Default()

	api := engine.Group("/api/v1")
	{
		product.RegisterRoutes(api, h.ProductHandler)
	}

	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return engine
}
