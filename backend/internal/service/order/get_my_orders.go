package order

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) GetMyOrders(ctx context.Context, userID string) (*[]model.Order, error) {
	orders, err := s.repo.GetMyOrders(ctx, userID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
