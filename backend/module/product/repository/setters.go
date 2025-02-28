package repository

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/module/product/converter"
	"github.com/Fi44er/sdmedik/backend/module/product/domain"
	"gorm.io/gorm"
)

func (r *ProductRepository) Create(ctx context.Context, productDomain *domain.Product, tx *gorm.DB) error {
	r.logger.Infof("Creating product: %+v...", productDomain)
	db := tx
	if db == nil {
		db = r.db
	}

	productModel := converter.ToModelFromDomain(productDomain)
	if err := db.WithContext(ctx).Create(productModel).Error; err != nil {
		r.logger.Errorf("Error creating product: %v", err)
		return err
	}

	productDomain.ID = productModel.ID

	r.logger.Info("Product created successfully")
	return nil
}
