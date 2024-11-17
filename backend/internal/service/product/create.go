package product

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) Create(ctx context.Context, product *model.Product) error {
	return s.productRepository.Create(ctx, product)
}
