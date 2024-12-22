package repository

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"gorm.io/gorm"
)

type ITransactionManager interface {
	BeginTransaction(ctx context.Context) (*gorm.DB, error)
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB)
	WithTransaction(tx *gorm.DB) *gorm.DB
}

type ICategoryRepository interface {
	Create(ctx context.Context, data *model.Category, tx *gorm.DB) error
	GetAll(ctx context.Context) ([]model.Category, error)
	GetByID(ctx context.Context, id int) (model.Category, error)
	Delete(ctx context.Context, id int) error
	GetByIDs(ctx context.Context, ids []int) ([]model.Category, error)
	GetByName(ctx context.Context, name string) (model.Category, error)
}

type IProductRepository interface {
	Create(ctx context.Context, data *model.Product, tx *gorm.DB) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, data *model.Product) error

	Get(ctx context.Context, criteria dto.ProductSearchCriteria) ([]model.Product, error)
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
	CreateMany(ctx context.Context, data *[]model.Characteristic, tx *gorm.DB) error
	GetByID(ctx context.Context, id int) (model.Characteristic, error)
	GetByCategoryID(ctx context.Context, categoryID int) ([]model.Characteristic, error)
	Delete(ctx context.Context, id int) error
}

type ICharacteristicValueRepository interface {
	Create(ctx context.Context, data *model.CharacteristicValue) error
	CreateMany(ctx context.Context, data *[]model.CharacteristicValue, tx *gorm.DB) error
}

type IImageRepository interface {
	CreateMany(ctx context.Context, data *[]model.Image, tx *gorm.DB) error
	GetByID(ctx context.Context, productID *string, categoryID *int, tx *gorm.DB) ([]model.Image, error)
}
