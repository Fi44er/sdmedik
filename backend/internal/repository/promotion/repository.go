package promotion

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	def "github.com/Fi44er/sdmedik/backend/internal/repository"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

var _ def.IPromotionRepository = (*repository)(nil)

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

func (r *repository) Create(ctx context.Context, data *model.Promotion) error {
	r.logger.Info("Creating promotion...")
	if err := r.db.WithContext(ctx).Create(data).Error; err != nil {
		r.logger.Errorf("Failed to create promotion: %v", err)
		return err
	}
	r.logger.Info("Promotion created successfully")
	return nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	r.logger.Infof("Deleting promotion with ID: %s...", id)
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Promotion{})
	if err := result.Error; err != nil {
		r.logger.Errorf("Failed to delete promotion: %v", err)
		return err
	}

	// if result.RowsAffected == 0 {
	// 	r.logger.Warnf("Promotion with ID %s not found", id)
	// 	return constants.ErrPromotionNotFound
	// }
	//
	r.logger.Infof("Promotion deleted by ID: %v successfully", id)
	return nil
}

func (r *repository) GetByID(ctx context.Context, id string) (*model.Promotion, error) {
	r.logger.Infof("Fetching promotion with ID: %s...", id)
	promotion := new(model.Promotion)
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(promotion).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Warnf("Promotion with ID %s not found", id)
			return nil, nil
		}
		r.logger.Errorf("Failed to fetch promotion with ID %s: %v", id, err)
		return nil, err
	}
	r.logger.Info("Promotion fetched successfully")
	return promotion, nil
}

func (r *repository) GetAll(ctx context.Context) (*[]model.Promotion, error) {
	r.logger.Info("Fetching promotions...")
	promotions := new([]model.Promotion)
	if err := r.db.WithContext(ctx).Find(promotions).Error; err != nil {
		r.logger.Errorf("Failed to fetch promotions: %v", err)
		return nil, err
	}
	r.logger.Info("Promotions fetched successfully")
	return promotions, nil
}
