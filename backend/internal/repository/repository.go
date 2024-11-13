package repository

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

type IProductRepository interface {
	Create(ctx context.Context, data *model.Product) error
	GetAll(ctx context.Context) ([]*model.Product, error)
	GetByID(ctx context.Context, id string) (*model.Product, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, data *model.Product) error
}
