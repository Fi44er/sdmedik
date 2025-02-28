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
