package service

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
)

type IUserService interface {
	Register(ctx context.Context, user *dto.Register) error
	Login(ctx context.Context, user *dto.Login) (string, string, error)
	Logout(ctx context.Context, refreshToken string, accessTokenUUID string) error
	RefreshAccessToken(ctx context.Context, refreshToken string) (string, error)

	GetByID(ctx context.Context, id string) (model.User, error)
	GetAll(ctx context.Context, offset int, limit int) ([]model.User, error)
	Update(ctx context.Context, data *dto.UpdateUser, id string) error
}

type IProductService interface {
	Create(ctx context.Context, product *model.Product) error
	GetAll(ctx context.Context, offset int, limit int) ([]model.Product, error)
	GetByID(ctx context.Context, id string) (model.Product, error)
	Update(ctx context.Context, product *model.Product) error
	Delete(ctx context.Context, id string) error
}
