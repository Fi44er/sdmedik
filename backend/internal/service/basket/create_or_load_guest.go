package basket

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) GetOrCreateGuestBasket(ctx context.Context, guestBasketID string) (*model.GuestBasket, error) {
	basket := new(model.GuestBasket)
	var err error
	if guestBasketID != "" {
		basket, err = s.repo.GetGuestBasketByID(ctx, guestBasketID)
		if err != nil {
			return nil, err
		}
	}

	if basket.ID != "" {
		if err := s.repo.CreateGuestBasket(ctx, basket); err != nil {
			return nil, err
		}
	}

	return basket, nil
}
