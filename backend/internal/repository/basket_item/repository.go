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

func (r *repository) Delete(ctx context.Context, id string) error {
	r.logger.Info("Deleting basket item...")
	result := r.db.WithContext(ctx).Delete(&model.BasketItem{}, id)
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
