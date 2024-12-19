package image

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	def "github.com/Fi44er/sdmedik/backend/internal/repository"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

var _ def.IImageRepository = (*repository)(nil)

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

func (r *repository) CreateMany(ctx context.Context, data *[]model.Image, tx *gorm.DB) error {
	r.logger.Info("Creating images...")
	db := tx
	if db == nil {
		r.logger.Error("Transaction is nil")
		db = r.db
	}

	if err := db.WithContext(ctx).Create(data).Error; err != nil {
		r.logger.Errorf("Failed to create images: %v", err)
		return err
	}

	r.logger.Infof("Images created successfully")
	return nil
}

func (r *repository) GetByProductID(ctx context.Context, productID string, tx *gorm.DB) ([]model.Image, error) {
	r.logger.Info("Getting images...")
	db := tx
	if db == nil {
		r.logger.Error("Transaction is nil")
		db = r.db
	}
	var images []model.Image
	if err := db.WithContext(ctx).Where("product_id = ?", productID).Find(&images).Error; err != nil {
		r.logger.Errorf("Failed to get images: %v", err)
		return nil, err
	}
	return images, nil
}
