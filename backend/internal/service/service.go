package service

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/internal/response"
	"gorm.io/gorm"
)

type IUserService interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetAll(ctx context.Context, offset int, limit int) (*[]model.User, error)
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
	Create(ctx context.Context, product *dto.CreateProduct, images *dto.Images) error
	Update(ctx context.Context, product *dto.UpdateProduct, images *dto.Images, id string) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, criteria dto.ProductSearchCriteria) (*[]model.Product, error)
	GetFilter(ctx context.Context, categoryID int) (*response.ProductFilter, error)
}

type ICategoryService interface {
	Create(ctx context.Context, data *dto.CreateCategory, image *dto.Image) error
	GetAll(ctx context.Context) (*[]model.Category, error)
	GetByID(ctx context.Context, id int) (*model.Category, error)
	Delete(ctx context.Context, id int) error
	GetByIDs(ctx context.Context, ids []int) (*[]model.Category, error)
}

type ICharacteristicService interface {
	Create(ctx context.Context, characteristic *dto.CreateCharacteristic) error
	CreateMany(ctx context.Context, characteristics *[]model.Characteristic, tx *gorm.DB) error
	GetByID(ctx context.Context, id int) (*model.Characteristic, error)
	GetByCategoryID(ctx context.Context, categoryID int) (*[]model.Characteristic, error)
	Delete(ctx context.Context, id int) error
}

type ICharacteristicValueService interface {
	Create(ctx context.Context, characteristicValue *dto.CharacteristicValue) error
	CreateMany(ctx context.Context, characteristicValues *[]model.CharacteristicValue, tx *gorm.DB) error
	DeleteByProductID(ctx context.Context, productID string, tx *gorm.DB) error
}

type IImageService interface {
	CreateMany(ctx context.Context, dto *dto.CreateImages, tx *gorm.DB) error
	DeleteByNames(ctx context.Context, names []string) error
	DeleteByIDs(ctx context.Context, ids []string, names []string, tx *gorm.DB) error
}

type IBasketService interface{}

type ISearchService interface {
	Search(ctx context.Context, query string) (*[]response.SearchRes, error)
	// AddProductToIndex(index bleve.Index, product *model.Product) error
}
