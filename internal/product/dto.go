package product

type CreateProductRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=255"`
	Description string `json:"description"`
	Price       int64  `json:"price" binding:"required,gte=0"`
	Stock       int    `json:"stock" binding:"required,gte=0"`
	IsActive    *bool  `json:"is_active"`
}

type UpdateProductRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Price       *int64  `json:"price" binding:"omitempty,gte=0"`
	Stock       *int    `json:"stock" binding:"omitempty,gte=0"`
	IsActive    *bool   `json:"is_active"`
}

type ProductResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int    `json:"stock"`
	IsActive    bool   `json:"is_active"`
}
