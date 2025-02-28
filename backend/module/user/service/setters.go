package service

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/module/user/domain"
)

func (s *UserService) Create(ctx context.Context, user *domain.User) error {
	if err := s.repo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) Update(ctx context.Context, user *domain.User) error {
	if err := s.repo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
