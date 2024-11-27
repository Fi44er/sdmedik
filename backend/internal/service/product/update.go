package product

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) Update(ctx context.Context, product *model.Product) error {
	return s.repo.Update(ctx, product)
}
