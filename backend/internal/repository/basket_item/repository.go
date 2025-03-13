package basketitem

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	def "github.com/Fi44er/sdmedik/backend/internal/repository"
	"github.com/Fi44er/sdmedik/backend/pkg/constants"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

var _ def.IBasketItemRepository = (*repository)(nil)

type repository struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewRepository(logger *logger.Logger, db *gorm.DB) *repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) Create(ctx context.Context, data *model.BasketItem) error {
	r.logger.Info("Creating basket item...")
	if err := r.db.WithContext(ctx).Create(data).Error; err != nil {
		r.logger.Errorf("Failed to create basket item: %v", err)
		return err
	}
	r.logger.Infof("Basket item created successfully")
	return nil
}

func (r *repository) Update(ctx context.Context, data *model.BasketItem) error {
	r.logger.Info("Updating basket item...")
	result := r.db.WithContext(ctx).Model(data).Updates(data)
	if err := result.Error; err != nil {
		r.logger.Errorf("Failed to update basket item: %v", err)
		return err
	}

	if result.RowsAffected == 0 {
		r.logger.Warnf("Basket item with ID %s not found", data.ID)
		return constants.ErrBasketItemNotFound
	}

	r.logger.Infof("Basket item updated successfully")
	return nil
}

func (r *repository) Delete(ctx context.Context, id string, basketID string) error {
	r.logger.Info("Deleting basket item...")
	result := r.db.WithContext(ctx).Where("id = ? AND basket_id = ?", id, basketID).Delete(&model.BasketItem{})
	if err := result.Error; err != nil {
		r.logger.Errorf("Failed to delete basket item: %v", err)
		return err
	}

	if result.RowsAffected == 0 {
		r.logger.Warnf("Basket item with ID %s not found", id)
		return constants.ErrBasketItemNotFound
	}

	r.logger.Infof("Basket item deleted successfully")
	return nil
}

func (r *repository) GetByProductBasketID(ctx context.Context, productID string, basketID string) (*model.BasketItem, error) {
	r.logger.Info("Fetching basket item by product and basket ID...")
	basketItem := new(model.BasketItem)
	if err := r.db.WithContext(ctx).Where("product_id = ? AND basket_id = ?", productID, basketID).First(basketItem).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		r.logger.Errorf("Failed to fetch basket item by product and basket ID: %v", err)
		return nil, err
	}
	r.logger.Info("Basket item fetched by product and basket ID successfully")
	return basketItem, nil
}

func (r *repository) UpdateItemQuantity(ctx context.Context, data *model.BasketItem) error {
	r.logger.Info("Updating basket item quantity...")
	result := r.db.WithContext(ctx).Model(data).Where("product_id = ? AND basket_id = ?", data.ProductID, data.BasketID).Updates(data)
	if err := result.Error; err != nil {
		r.logger.Errorf("Failed to update basket item quantity: %v", err)
		return err
	}

	if result.RowsAffected == 0 {
		r.logger.Warnf("Basket item with ID %s not found", data.ID)
		return constants.ErrBasketItemNotFound
	}

	r.logger.Infof("Basket item quantity updated successfully")
	return nil
}

func (r *repository) GetByProductIDIsoIsCert(ctx context.Context, productID string, basketID string, iso string, isCert bool) (*model.BasketItem, error) {
	r.logger.Info("Fetching basket item by product and basket ID...")
	basketItem := new(model.BasketItem)
	if err := r.db.WithContext(ctx).Where("product_id = ? AND basket_id = ? AND iso = ? AND is_certificate = ?", productID, basketID, iso, isCert).First(basketItem).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		r.logger.Errorf("Failed to fetch basket item by product and basket ID: %v", err)
		return nil, err
	}
	r.logger.Info("Basket item fetched by product and basket ID successfully")
	return basketItem, nil
}
