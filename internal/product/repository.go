package product

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func ProductRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, p *Product) error {
	query := `
		INSERT INTO products
			(name, description, price, stock, is_active)
		VALUES
			($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowxContext(
		ctx,
		query,
		p.Name,
		p.Description,
		p.Price,
		p.Stock,
		p.IsActive,
	).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
}

func (r *Repository) Get(ctx context.Context, id string) (*Product, error) {
	query := `
		SELECT id, name, description, price, stock, is_active, created_at, updated_at
		FROM products
		WHERE id = $1
	`

	p := &Product{}
	err := r.db.QueryRowxContext(ctx, query, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.IsActive, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return p, nil
}
