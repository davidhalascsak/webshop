package router

import (
	"net/http"
	"webshop/internal/api"
	"webshop/internal/product"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	ProductHandler *product.Handler
}

func Setup(h Handlers) *gin.Engine {
	engine := gin.Default()
	engine.Use(api.ErrorMiddleware())

	apiGroup := engine.Group("/api/v1")
	{
		product.RegisterRoutes(apiGroup, h.ProductHandler)
	}

	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return engine
}
