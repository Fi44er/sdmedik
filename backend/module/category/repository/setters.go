package repository

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/module/category/converter"
	"github.com/Fi44er/sdmedik/backend/module/category/domain"
	"gorm.io/gorm"
)

func (r *CategoryRepository) Create(ctx context.Context, categoryDomain *domain.Category, tx *gorm.DB) error {
	r.logger.Infof("Creating category: %+v...", categoryDomain)
	db := tx
	if db == nil {
		db = r.db
	}

	categoryModel := converter.ToModelFromDomain(categoryDomain)
	if err := db.WithContext(ctx).Create(categoryModel).Error; err != nil {
		r.logger.Errorf("Error creating category: %v", err)
		return err
	}

	categoryDomain.ID = categoryModel.ID

	r.logger.Info("Category created successfully")
	return nil
}
