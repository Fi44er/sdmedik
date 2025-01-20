package product

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) GetByIDs(ctx context.Context, ids []string) (*[]model.Product, error) {
	return s.repo.GetByIDs(ctx, ids)
}
