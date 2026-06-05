package product

import "context"

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
	return s.repo.Get(ctx, id)
}
