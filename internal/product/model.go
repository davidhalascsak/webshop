package product

import (
	"time"
)

type Product struct {
	ID          uint      `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Price       int64     `db:"price" json:"price"`
	Stock       int       `db:"stock" json:"stock"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
