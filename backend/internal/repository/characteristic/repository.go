package characteristic

import (
	"context"
	"fmt"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	def "github.com/Fi44er/sdmedik/backend/internal/repository"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

var _ def.ICharacteristicRepository = (*repository)(nil)

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

func (r *repository) Create(ctx context.Context, data *model.Characteristic) error {
	r.logger.Info("Creating characteristic...")
	if err := r.db.WithContext(ctx).Create(data).Error; err != nil {
		r.logger.Errorf("Failed to create characteristic: %v", err)
		return err
	}

	r.logger.Infof("Characteristic created successfully")
	return nil
}

func (r *repository) GetByID(ctx context.Context, id int) (model.Characteristic, error) {
	r.logger.Infof("Fetching characteristic with ID: %v...", id)
	var characteristic model.Characteristic
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&characteristic).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Warnf("Characteristic with ID %v not found", id)
			return characteristic, nil
		}
		r.logger.Errorf("Failed to fetch characteristic with ID %v: %v", id, err)
		return model.Characteristic{}, err
	}
	r.logger.Info("Characteristic fetched successfully")
	return characteristic, nil
}

func (r *repository) GetByCategoryID(ctx context.Context, categoryID string) ([]model.Characteristic, error) {
	r.logger.Infof("Fetching characteristic with category ID: %s...", categoryID)
	var characteristics []model.Characteristic
	if err := r.db.WithContext(ctx).Where("category_id = ?", categoryID).Find(&characteristics).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Warnf("Characteristic with category ID %s not found", categoryID)
			return characteristics, nil
		}
		r.logger.Errorf("Failed to fetch characteristic with category ID %s: %v", categoryID, err)
		return nil, err
	}
	r.logger.Info("Characteristic fetched successfully")
	return characteristics, nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	r.logger.Infof("Deleting characteristic with ID: %s...", id)
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Characteristic{})
	if err := result.Error; err != nil {
		r.logger.Errorf("Failed to delete characteristic: %v", err)
		return err
	}

	if result.RowsAffected == 0 {
		r.logger.Warnf("Characteristic with ID %s not found", id)
		return fmt.Errorf("Characteristic not found")
	}

	r.logger.Infof("Characteristic deleted by ID: %v successfully", id)
	return nil
}
