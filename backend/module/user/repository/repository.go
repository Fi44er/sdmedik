package repository

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/module/user/domain"
	"github.com/Fi44er/sdmedik/backend/shared/logger"
	"gorm.io/gorm"
)

var _ IUserRepository = (*UserRepository)(nil)

type IUserRepository interface {
	Create(ctx context.Context, userDomain *domain.User) error
	Update(ctx context.Context, userDomain *domain.User) error
	Delete(ctx context.Context, id string) error

	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	GetAll(ctx context.Context, offset, limit int) ([]domain.User, error)
}

type UserRepository struct {
	logger *logger.Logger
	db     *gorm.DB
}

func NewUserRepository(logger *logger.Logger, db *gorm.DB) *UserRepository {
	return &UserRepository{
		logger: logger,
		db:     db,
	}
}
