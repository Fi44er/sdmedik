package repository

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

type ICategoryRepository interface {
	Create(ctx context.Context, data *model.Category) error
	GetAll(ctx context.Context) ([]*model.Category, error)
	GetByID(ctx context.Context, id int) (*model.Category, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, data *model.Category) error
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
	GetAll(ctx context.Context, offset int, limit int) ([]model.Product, error)
	GetByID(ctx context.Context, id string) (model.Product, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, data *model.Product) error
}

type IUserRepository interface {
	Create(ctx context.Context, data *model.User) error
	GetAll(ctx context.Context, offset int, limit int) ([]model.User, error)
	GetByID(ctx context.Context, id string) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, data *model.User) error
}
