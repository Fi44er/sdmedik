package service

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/module/product/domain"
	customerr "github.com/Fi44er/sdmedik/backend/shared/custom_err"
)

func (s *ProductService) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		s.logger.Infof("Product with id %s not found", id)
		return nil, customerr.ErrProductNotFound
	}

	return product, nil
}

func (s *ProductService) GetAll(ctx context.Context) ([]domain.Product, error) {
	products, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		s.logger.Info("No products found")
		return nil, customerr.ErrProductNotFound
	}

	return products, nil
}
