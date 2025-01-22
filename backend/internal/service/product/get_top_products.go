package product

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/response"
)

func (s *service) GetTopProducts(ctx context.Context, limit int) ([]response.ProductPopularity, error) {
	return s.repo.GetTopProducts(ctx, limit)
}
