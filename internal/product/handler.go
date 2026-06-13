package product

import (
	"net/http"

	"webshop/internal/api"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func ProductHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	p, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	resp := ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		IsActive:    p.IsActive,
	}

	api.SuccessResponse(c, resp)
}

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")

	p, err := h.service.Get(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	resp := ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		IsActive:    p.IsActive,
	}

	api.SuccessResponse(c, resp)
}

func (h *Handler) Update(c *gin.Context) {
	// TODO: implement update handler
}

func (h *Handler) Delete(c *gin.Context) {
	// TODO: implement delete handler
}
