package user

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) GetByEmail(ctx context.Context, email string) (model.User, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
