package service

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/blevesearch/bleve/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type IUserService interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetAll(ctx context.Context, offset int, limit int) (*response.UsersResponse, error)
	Update(ctx context.Context, data *dto.UpdateUser, id string) error
	Delete(ctx context.Context, id string) error
}

type IAuthService interface {
	Register(ctx context.Context, user *dto.Register) error
	Login(ctx context.Context, user *dto.Login, userAgent string, sessionRes *session.Session) (string, string, error)
	Logout(ctx context.Context, refreshToken string, accessTokenUUID string, userAgent string) error
	RefreshAccessToken(ctx context.Context, refreshToken string, userAgent string) (string, error)
	VerifyCode(ctx context.Context, data *dto.VerifyCode) error
	SendCode(ctx context.Context, email string) error
}

type IProductService interface {
	Create(ctx context.Context, product *dto.CreateProduct, images *dto.Images) error
	CreateMany(ctx context.Context, products *[]dto.CreateProduct) error
	Update(ctx context.Context, product *dto.UpdateProduct, images *dto.Images, id string) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, criteria dto.ProductSearchCriteria) (*[]response.ProductResponse, *int64, error)
	GetFilter(ctx context.Context, categoryID int) (*response.ProductFilter, error)
	GetByIDs(ctx context.Context, ids []string) (*[]model.Product, error)
	GetTopProducts(ctx context.Context, limit int) (*[]response.TopProductRes, error)
}

type ICategoryService interface {
	Create(ctx context.Context, data *dto.CreateCategory, image *dto.Image) error
	GetAll(ctx context.Context) (*[]model.Category, error)
	GetByID(ctx context.Context, id int) (*model.Category, error)
	Delete(ctx context.Context, id int) error
	GetByIDs(ctx context.Context, ids []int) (*[]model.Category, error)
	Update(ctx context.Context, categoryID int, data *dto.UpdateCategory) error
}

type ICharacteristicService interface {
	Create(ctx context.Context, characteristic *dto.CreateCharacteristic) error
	CreateMany(ctx context.Context, characteristics *[]model.Characteristic, tx *gorm.DB) error
	GetByID(ctx context.Context, id int) (*model.Characteristic, error)
	GetByCategoryID(ctx context.Context, categoryID int) (*[]model.Characteristic, error)
	Delete(ctx context.Context, id int) error
	GetByIDs(ctx context.Context, ids []int) (*[]model.Characteristic, error)
	Update(ctx context.Context, categoryID int, characteristics []dto.UpdateCharacteristic, tx *gorm.DB) error
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

type IBasketService interface {
	Create(ctx context.Context, dto *dto.CreateBasket) error
	AddItem(ctx context.Context, dto *dto.AddBasketItem, userID string, sess *session.Session) error
	DeleteItem(ctx context.Context, itemID string, userID string, sess *session.Session) error
	GetByUserID(ctx context.Context, userID string, sess *session.Session) (*response.BasketResponse, error)
}

type ISearchService interface {
	Search(ctx context.Context, query string) (*[]response.SearchRes, error)
}

type IIndexService interface {
	Get() bleve.Index
}

type IWebScraperService interface {
	Scraper() error
}

type ICertificateService interface {
	CreateMany(ctx context.Context, data *[]model.Certificate) error
	UpdateMany(ctx context.Context, data *[]model.Certificate) error
	GetMany(ctx context.Context, data *[]dto.GetManyCert) (*[]model.Certificate, error)
}

type IOrderService interface {
	Create(ctx context.Context, data *dto.CreateOrder, userID string, sess *session.Session) (string, error)
	NotAuthCreate(ctx context.Context, data *dto.CreateOrder, productID string) (string, error)

	ChangeStatus(ctx context.Context, data *dto.ChangeOrderStatus) error
	// Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context, offset int, limit int) (*[]model.Order, error)
	GetMyOrders(ctx context.Context, userID string) (*[]model.Order, error)
}

type IPromotionService interface {
	Create(ctx context.Context, data *dto.CreatePromotion) error
	CheckAndApplyPromotions(ctx context.Context, basket *response.BasketResponse) (*response.BasketResponse, error)
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) (*[]model.Promotion, error)
}
