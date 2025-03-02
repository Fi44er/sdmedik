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

func (r *ProductRepository) CreateProductCategoies(ctx context.Context, productCategory []domain.ProductCategory, tx *gorm.DB) error {
	r.logger.Infof("Creating product categories: %+v...", productCategory)
	db := tx
	if db == nil {
		db = r.db
	}

	productCategoryModel := converter.ToModelProductCategorySliceFromDomain(productCategory)
	if err := db.WithContext(ctx).Create(productCategoryModel).Error; err != nil {
		r.logger.Errorf("Error creating product categories: %v", err)
		return err
	}

	r.logger.Info("Product categories created successfully")
	return nil
}
