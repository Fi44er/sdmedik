package user

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func (s *service) GetAll(ctx context.Context, offset int, limit int) ([]model.User, error) {
	users, err := s.repo.GetAll(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New(404, "Users not found")
	}

	return users, nil
}
