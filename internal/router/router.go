package router

import (
	"net/http"
	"webshop/internal/api"
	"webshop/internal/auth"
	"webshop/internal/product"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	ProductHandler *product.Handler
}

func Setup(h Handlers, authenticator *auth.Authenticator) *gin.Engine {
	engine := gin.Default()
	engine.Use(api.ErrorMiddleware())

	apiGroup := engine.Group("/api/v1")
	apiGroup.Use(api.AuthMiddleware(authenticator))
	{
		product.RegisterRoutes(apiGroup, h.ProductHandler)
	}

	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return engine
}
