package repository

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

type ICategoryRepository interface {
	Create(ctx context.Context, data *model.Category) error
	GetAll(ctx context.Context) ([]model.Category, error)
	GetByID(ctx context.Context, id int) (model.Category, error)
	Delete(ctx context.Context, id int) error
	GetByIDs(ctx context.Context, ids []int) ([]model.Category, error)
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

type ICharacteristicRepository interface {
	Create(ctx context.Context, data *model.Characteristic) error
	GetByID(ctx context.Context, id int) (model.Characteristic, error)
	GetByCategoryID(ctx context.Context, categoryID string) ([]model.Characteristic, error)
	Delete(ctx context.Context, id string) error
}
