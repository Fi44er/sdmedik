package service

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"gorm.io/gorm"
)

type IUserService interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id string) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
	GetAll(ctx context.Context, offset int, limit int) ([]model.User, error)
	Update(ctx context.Context, data *dto.UpdateUser, id string) error
	Delete(ctx context.Context, id string) error
}

type IAuthService interface {
	Register(ctx context.Context, user *dto.Register) error
	Login(ctx context.Context, user *dto.Login) (string, string, error)
	Logout(ctx context.Context, refreshToken string, accessTokenUUID string) error
	RefreshAccessToken(ctx context.Context, refreshToken string) (string, error)
	VerifyCode(ctx context.Context, data *dto.VerifyCode) error
	SendCode(ctx context.Context, email string) error
}

type IProductService interface {
	Create(ctx context.Context, product *dto.CreateProduct) error
	GetAll(ctx context.Context, offset int, limit int) ([]model.Product, error)
	GetByID(ctx context.Context, id string) (model.Product, error)
	Update(ctx context.Context, product *model.Product) error
	Delete(ctx context.Context, id string) error
}

type ICategoryService interface {
	Create(ctx context.Context, data *dto.CreateCategory) error
	GetAll(ctx context.Context) ([]model.Category, error)
	GetByID(ctx context.Context, id int) (model.Category, error)
	Delete(ctx context.Context, id int) error
	GetByIDs(ctx context.Context, ids []int) ([]model.Category, error)
}

type ICharacteristicService interface {
	Create(ctx context.Context, characteristic *dto.CreateCharacteristic) error
	CreateMany(ctx context.Context, characteristics *[]model.Characteristic, tx *gorm.DB) error
	GetByID(ctx context.Context, id int) (model.Characteristic, error)
	GetByCategoryID(ctx context.Context, categoryID int) ([]model.Characteristic, error)
	Delete(ctx context.Context, id int) error
}

type ICharacteristicValueService interface {
	Create(ctx context.Context, characteristicValue *dto.CharacteristicValue) error
	CreateMany(ctx context.Context, characteristicValues *[]model.CharacteristicValue, tx *gorm.DB) error
}
