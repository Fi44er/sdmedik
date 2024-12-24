package characteristicvalue

import (
	"context"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	def "github.com/Fi44er/sdmedik/backend/internal/repository"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

var _ def.ICharacteristicValueRepository = (*repository)(nil)

type repository struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewRepository(
	logger *logger.Logger,
	db *gorm.DB,
) *repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) Create(ctx context.Context, data *model.CharacteristicValue) error {
	r.logger.Info("Creating characteristic value...")
	if err := r.db.WithContext(ctx).Create(data).Error; err != nil {
		r.logger.Errorf("Failed to create characteristic value: %v", err)
		return err
	}

	r.logger.Infof("Characteristic value created successfully")
	return nil
}

func (r *repository) CreateMany(ctx context.Context, data *[]model.CharacteristicValue, tx *gorm.DB) error {
	r.logger.Info("Creating characteristic values...")
	db := tx
	if db == nil {
		r.logger.Error("Transaction is nil")
		db = r.db
	}

	if err := db.WithContext(ctx).Create(data).Error; err != nil {
		r.logger.Errorf("Failed to create characteristic values: %v", err)
		return err
	}

	r.logger.Infof("Characteristic values created successfully")
	return nil
}

func (r *repository) DeleteByProductID(ctx context.Context, productID string, tx *gorm.DB) error {
	r.logger.Info("Deleting characteristic values...")
	db := tx
	if db == nil {
		r.logger.Error("Transaction is nil")
		db = r.db
	}

	if err := db.WithContext(ctx).Where("product_id = ?", productID).Delete(&model.CharacteristicValue{}).Error; err != nil {
		r.logger.Errorf("Failed to delete characteristic values: %v", err)
		return err
	}

	r.logger.Infof("Characteristic values deleted successfully")
	return nil
}
