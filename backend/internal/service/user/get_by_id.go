package user

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func (s *service) GetByID(ctx context.Context, id string) (model.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return model.User{}, err
	}

	if user.ID == "" {
		return model.User{}, errors.New(404, "User not found")
	}

	return user, nil
}