package product

import (
	"webshop/internal/api"
	"webshop/internal/auth"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(apiGroup *gin.RouterGroup, h *Handler) {
	products := apiGroup.Group("/products")
	{
		products.POST("", api.RequireRole(auth.RoleAdmin), h.Create)
		products.GET("/:id", api.RequireRole(auth.RoleUser, auth.RoleAdmin), h.Get)
	}
}
