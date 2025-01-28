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

func (r *repository) GetByID(ctx context.Context, productID *string, categoryID *int, tx *gorm.DB) (*[]model.Image, error) {
	r.logger.Info("Getting images...")
	db := tx
	if db == nil {
		r.logger.Error("Transaction is nil")
		db = r.db
	}
	images := new([]model.Image)
	request := db.WithContext(ctx)

	// Проверяем, какой идентификатор передан
	if productID != nil {
		request = request.Where("product_id = ?", productID)
	} else if categoryID != nil {
		request = request.Where("category_id = ?", categoryID)
	}

	if err := request.Find(images).Error; err != nil {
		r.logger.Errorf("Failed to get images: %v", err)
		return nil, err
	}
	return images, nil
}

func (r *repository) DeleteByIDs(ctx context.Context, id []string, tx *gorm.DB) error {
	r.logger.Info("Deleting image...")
	db := tx
	if db == nil {
		r.logger.Error("Transaction is nil")
		db = r.db
	}

	if err := db.WithContext(ctx).Where("id IN (?)", id).Delete(&model.Image{}).Error; err != nil {
		r.logger.Errorf("Failed to delete image: %v", err)
		return err
	}

	r.logger.Infof("Image deleted successfully")
	return nil
}
