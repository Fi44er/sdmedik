package repository

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/internal/response"
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
	GetAll(ctx context.Context) (*[]model.Category, error)
	GetByID(ctx context.Context, id int) (*model.Category, error)
	Delete(ctx context.Context, id int, tx *gorm.DB) error
	GetByIDs(ctx context.Context, ids []int) (*[]model.Category, error)
	GetByName(ctx context.Context, name string) (*model.Category, error)
	Update(ctx context.Context, cateegory *model.Category, tx *gorm.DB) error
}

type IProductRepository interface {
	Create(ctx context.Context, data *model.Product, tx *gorm.DB) error
	CreateMany(ctx context.Context, data *[]model.Product) error

	Delete(ctx context.Context, id string, tx *gorm.DB) error
	Update(ctx context.Context, data *model.Product, tx *gorm.DB) error
	DeleteCategoryAssociation(ctx context.Context, productID string, tx *gorm.DB) error

	Get(ctx context.Context, criteria dto.ProductSearchCriteria) (*[]model.Product, *int64, error)
	GetByIDs(ctx context.Context, ids []string) (*[]model.Product, error)
	GetTopProducts(ctx context.Context, limit int) ([]response.ProductPopularity, error)

	GetByArticles(ctx context.Context, articles []string) (*[]model.Product, error)
}

type IUserRepository interface {
	Create(ctx context.Context, data *model.User) error
	GetAll(ctx context.Context, offset int, limit int) (*[]model.User, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, data *model.User) error
}

type ICharacteristicRepository interface {
	Create(ctx context.Context, data *model.Characteristic) error
	CreateMany(ctx context.Context, data *[]model.Characteristic, tx *gorm.DB) error
	GetByID(ctx context.Context, id int) (*model.Characteristic, error)
	GetByCategoryID(ctx context.Context, categoryID int) (*[]model.Characteristic, error)
	Delete(ctx context.Context, id int) error
	GetByIDs(ctx context.Context, ids []int) (*[]model.Characteristic, error)
	Update(ctx context.Context, data *model.Characteristic, tx *gorm.DB) error
	DeleteMany(ctx context.Context, ids []int, tx *gorm.DB) error
}

type ICharacteristicValueRepository interface {
	Create(ctx context.Context, data *model.CharacteristicValue) error
	CreateMany(ctx context.Context, data *[]model.CharacteristicValue, tx *gorm.DB) error
	DeleteByProductID(ctx context.Context, productID string, tx *gorm.DB) error
}

type IImageRepository interface {
	CreateMany(ctx context.Context, data *[]model.Image, tx *gorm.DB) error
	GetByID(ctx context.Context, productID *string, categoryID *int, tx *gorm.DB) (*[]model.Image, error)
	DeleteByIDs(ctx context.Context, id []string, tx *gorm.DB) error
}

type IBasketRepository interface {
	Create(ctx context.Context, data *model.Basket) error
	GetByUserID(ctx context.Context, userID string) (*model.Basket, error)
}

type IBasketItemRepository interface {
	Create(ctx context.Context, data *model.BasketItem) error
	Update(ctx context.Context, data *model.BasketItem) error
	Delete(ctx context.Context, itemID string, basketID string) error
	GetByProductBasketID(ctx context.Context, productID string, basketID string) (*model.BasketItem, error)
	UpdateItemQuantity(ctx context.Context, dto *model.BasketItem) error
	GetByProductIDIsoIsCert(ctx context.Context, productID string, basketID string, iso string, isCert bool) (*model.BasketItem, error)
}

type ICertificateRepository interface {
	CreateMany(ctx context.Context, data *[]model.Certificate) error
	UpdateMany(ctx context.Context, data *[]model.Certificate) error
	GetMany(ctx context.Context, data *[]dto.GetManyCert) (*[]model.Certificate, error)
}

type IOrderRepository interface {
	Create(ctx context.Context, data *model.Order) error
	AddItems(ctx context.Context, items *[]model.OrderItem) error
	GetAll(ctx context.Context, offset int, limit int) (*[]model.Order, error)
	GetMyOrders(ctx context.Context, userID string) (*[]model.Order, error)
	Update(ctx context.Context, data *model.Order) error
}

type IPromotionRepository interface {
	Create(ctx context.Context, data *model.Promotion) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*model.Promotion, error)
	GetAll(ctx context.Context) (*[]model.Promotion, error)

	CreateConditions(ctx context.Context, condition *model.Condition) error
	CreateRewards(ctx context.Context, reward *model.Reward) error
}

type IChatRepository interface {
	GetByID(ctx context.Context, id string) (*model.Chat, error)
	Create(ctx context.Context, chat *model.Chat) error
	GetAll(ctx context.Context, offset, limit int) ([]model.Chat, error)
	DeleteChat(ctx context.Context, chatID string) error

	SaveMessage(ctx context.Context, message *model.Message) error
	GetMessagesByChatID(ctx context.Context, chatID string) ([]model.Message, error)

	MarkMsgAsRead(ctx context.Context, msgID string) error
	GetMessageByID(ctx context.Context, id string) (*model.Message, error)
	GetUnreadCount(ctx context.Context, chatID, userID string) (int64, error)

	CreateFragment(ctx context.Context, fragment *model.Fragment) error
	AddEndMsgID(ctx context.Context, chatID, msgID string) error
	GetFragmentByID(ctx context.Context, id string) (*model.Fragment, error)
	GetLastChatFragment(ctx context.Context, chatID string) (*model.Fragment, error)

	GetFragmentsByChatID(ctx context.Context, chatId string) ([]model.Fragment, error)
	GetMessagesInFragment(ctx context.Context, fragment model.Fragment) ([]model.Message, error)
	UpdateFragment(ctx context.Context, fragment *model.Fragment) error
	GetActiveFragment(ctx context.Context, chatID string) (*model.Fragment, error)
}

type IPageRepository interface {
	AddOrUpdateElement(ctx context.Context, element *model.Element) error
	GetByPath(ctx context.Context, path string) (*model.Page, error)
	Create(ctx context.Context, data *model.Page) error
}
