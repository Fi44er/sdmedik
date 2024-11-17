package repository

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

type ICategoryRepository interface {
	Create(ctx context.Context, data *model.Categori) error
	GetAll(ctx context.Context) ([]*model.Categori, error)
	GetByID(ctx context.Context, id int) (*model.Categori, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, data *model.Categori) error
}

type IImageRepository interface {
	Create(ctx context.Context, data *model.Image) error
	GetAll(ctx context.Context) ([]*model.Image, error)
	GetByID(ctx context.Context, id string) (*model.Image, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, data *model.Image) error
}

type IOrderRepository interface {
	Create(ctx context.Context, data *model.Order) error
	GetAll(ctx context.Context) ([]*model.Order, error)
	GetByID(ctx context.Context, id string) (*model.Order, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, data *model.Order) error
}

type IPriceRepository interface {
	Create(ctx context.Context, data *model.Price) error
	GetAll(ctx context.Context) ([]*model.Price, error)
	GetByID(ctx context.Context, id string) (*model.Price, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, data *model.Price) error
}

type IProductRepository interface {
	Create(ctx context.Context, data *model.Product) error
	GetAll(ctx context.Context) ([]model.Product, error)
	GetByID(ctx context.Context, id string) (*model.Product, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, data *model.Product) error
}

type ITokenRepository interface {
	Create(ctx context.Context, data *model.Token) error
	GetByUserID(ctx context.Context, id string) (*model.Token, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, data *model.Token) error
}

type IUserRepository interface {
	Create(ctx context.Context, data *model.User) error
	GetAll(ctx context.Context) ([]*model.User, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, data *model.User) error
}
