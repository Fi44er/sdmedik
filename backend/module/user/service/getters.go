package service

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/module/user/domain"
	"github.com/Fi44er/sdmedik/backend/shared/custom_err"
)

func (s *UserService) GetAll(ctx context.Context, offset, limit int) ([]domain.User, error) {
	users, err := s.repo.GetAll(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		s.logger.Infof("No users found")
		return nil, customerr.ErrUserNotFound
	}

	return users, nil
}

func (s *UserService) GetByID(ctx context.Context, id string) (*domain.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		s.logger.Infof("User with id %s not found", id)
		return nil, customerr.ErrUserNotFound
	}

	return user, nil
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		s.logger.Infof("User with email %s not found", email)
		return nil, customerr.ErrUserNotFound
	}

	return user, nil
}
