package order

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) GetAll(ctx context.Context, offset int, limit int) (*[]model.Order, error) {
	orders, err := s.repo.GetAll(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
