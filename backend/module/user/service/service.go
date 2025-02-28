package service

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/module/user/domain"
	"github.com/Fi44er/sdmedik/backend/module/user/repository"
	"github.com/Fi44er/sdmedik/backend/shared/logger"
)

var _ IUserService = (*UserService)(nil)

type IUserService interface {
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id string) error

	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	GetAll(ctx context.Context, offset, limit int) ([]domain.User, error)
}

type UserService struct {
	repo repository.IUserRepository

	logger *logger.Logger
}

func NewUserService(
	repo repository.IUserRepository,
	logger *logger.Logger,
) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}
