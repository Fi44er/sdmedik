package basket

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/constants"
)

func (s *service) GetByUserID(ctx context.Context, userID string) (*model.Basket, error) {
	basket, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if basket == nil {
		return nil, constants.ErrBasketNotFound
	}

	return basket, nil
}
