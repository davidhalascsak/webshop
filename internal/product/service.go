package product

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"webshop/internal/apperrors"
)

type Service struct {
	repo *Repository
}

func ProductService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, req CreateProductRequest) (*Product, error) {
	isActive := true
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	p := &Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		IsActive:    isActive,
	}

	if err := s.repo.Create(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *Service) Get(ctx context.Context, id string) (*Product, error) {
	p, err := s.repo.Get(ctx, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &apperrors.AppError{
				Kind:    apperrors.ErrNotFound,
				Message: fmt.Sprintf("Product with id %s is not found", id),
			}
		}
		return nil, err
	}

	return p, nil
}
