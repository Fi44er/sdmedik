package repository

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/module/product/converter"
	"github.com/Fi44er/sdmedik/backend/module/product/domain"
	"github.com/Fi44er/sdmedik/backend/module/product/model"
	"gorm.io/gorm"
)

func (r *ProductRepository) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	r.logger.Infof("Getting product by ID: %s...", id)
	var product model.Product
	err := r.db.First(&product, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Warnf("Product with ID %s not found", id)
			return nil, nil
		}
		r.logger.Errorf("Error getting product by ID %v", err)
		return nil, err
	}

	files, err := r.fileRepo.GetFilesByOwner(ctx, product.ID, "product")
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		product.ImageIDs = append(product.ImageIDs, f.ID)
	}

	productDomain := converter.ToDomainFromModel(&product)

	return productDomain, nil
}

func (r *ProductRepository) GetAll(ctx context.Context) ([]domain.Product, error) {
	r.logger.Info("Getting all products...")
	var products []model.Product
	err := r.db.Find(&products).Error
	if err != nil {
		r.logger.Errorf("Error getting all products: %v", err)
		return nil, err
	}

	productDomains := make([]domain.Product, len(products))
	for i, product := range products {
		files, err := r.fileRepo.GetFilesByOwner(ctx, product.ID, "product")
		if err != nil {
			return nil, err
		}
		for _, f := range files {
			product.ImageIDs = append(product.ImageIDs, f.ID)
		}

		productDomains[i] = *converter.ToDomainFromModel(&product)
	}

	return productDomains, nil
}

func (r *ProductRepository) GetByCategoryID(ctx context.Context, categoryID string) ([]domain.Product, error) {
	r.logger.Infof("Getting products by category ID: %s...", categoryID)

	var products []model.Product
	err := r.db.Preload("Categories").
		Joins("JOIN product_categories ON products.id = product_categories.product_id").
		Where("product_categories.category_id = ?", categoryID).
		Find(&products).Error
	if err != nil {
		return nil, err
	}

	productDomains := converter.ToDomainSlicceFromModel(products)

	return productDomains, nil
}
