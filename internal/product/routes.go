package product

import "github.com/gin-gonic/gin"

func RegisterRoutes(api *gin.RouterGroup, h *Handler) {
	products := api.Group("/products")
	{
		products.POST("/", h.Create)
	}
}
