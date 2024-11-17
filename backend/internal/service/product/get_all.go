package product

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func (s *service) GetAll(ctx context.Context, offset int, limit int) ([]model.Product, error) {
	products, err := s.productRepository.GetAll(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, errors.New(404, "no products found")
	}

	return products, nil
}
