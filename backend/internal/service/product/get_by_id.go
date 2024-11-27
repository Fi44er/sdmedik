package product

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func (s *service) GetByID(ctx context.Context, id string) (model.Product, error) {
	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return model.Product{}, err
	}

	if product.ID == "" {
		return model.Product{}, errors.New(404, "product not found")
	}

	return product, nil
}
