package user

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func (s *service) GetAll(ctx context.Context, offset int, limit int) (*response.UsersResponse, error) {
	users, err := s.repo.GetAll(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	if len(*users) == 0 {
		return nil, errors.New(404, "Users not found")
	}

	usersRes := response.UsersResponse{
		Users: make([]response.UserResponse, len(*users)),
		Count: len(*users),
	}

	for i, user := range *users {
		usersRes.Users[i] = response.FilterUserResponse(&user)
	}

	return &usersRes, nil
}
