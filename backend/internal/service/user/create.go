package user

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) Create(ctx context.Context, user *model.User) error {
	if err := s.repo.Create(ctx, user); err != nil {
		return err
	}

	if err := s.basketService.Create(ctx, &dto.CreateBasket{UserID: user.ID}); err != nil {
		return err
	}

	return nil
}
