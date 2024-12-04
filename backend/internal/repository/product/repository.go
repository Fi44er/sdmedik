package product

import (
	"context"
	"fmt"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	def "github.com/Fi44er/sdmedik/backend/internal/repository"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

var _ def.IProductRepository = (*repository)(nil)

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

func (r *repository) Create(ctx context.Context, data *model.Product) error {
	r.logger.Info("Creating product...")
	if err := r.db.WithContext(ctx).Create(data).Error; err != nil {
		r.logger.Errorf("Failed to create product: %v", err)
		return err
	}

	r.logger.Infof("Product created successfully")
	return nil
}

func (r *repository) GetByID(ctx context.Context, id string) (model.Product, error) {
	r.logger.Infof("Fetching product with ID: %s...", id)
	var product model.Product
	if err := r.db.WithContext(ctx).Preload("Categories").Where("id = ?", id).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Warnf("Product with ID %s not found", id)
			return product, nil
		}
		r.logger.Errorf("Failed to fetch product with ID %s: %v", id, err)
		return model.Product{}, err
	}
	r.logger.Info("Product fetched successfully")
	return product, nil
}

func (r *repository) GetAll(ctx context.Context, offset int, limit int) ([]model.Product, error) {
	r.logger.Info("Fetching productr...")
	var products []model.Product
	if offset == 0 {
		offset = -1
	}

	if limit == 0 {
		limit = -1
	}

	if err := r.db.WithContext(ctx).Preload("Categories").Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		r.logger.Errorf("Failed to fetch products: %v", err)
		return nil, err
	}
	r.logger.Info("Products fetched successfully")
	return products, nil
}

func (r *repository) Update(ctx context.Context, data *model.Product) error {
	r.logger.Info("Updating product...")

	result := r.db.WithContext(ctx).Model(data).Updates(data)
	if err := result.Error; err != nil {
		r.logger.Errorf("Failed to update product: %v", err)
		return err
	}

	if result.RowsAffected == 0 {
		r.logger.Warnf("Product with ID %s not found", data.ID)
		return fmt.Errorf("Product not found")
	}

	r.logger.Info("Product updated successfully")
	return nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	r.logger.Infof("Deleting product with ID: %s...", id)
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Product{})
	if err := result.Error; err != nil {
		r.logger.Errorf("Failed to delete product: %v", err)
		return err
	}

	if result.RowsAffected == 0 {
		r.logger.Warnf("Product with ID %s not found", id)
		return fmt.Errorf("Product not found")
	}

	r.logger.Infof("Product deleted by ID: %v successfully", id)
	return nil
}
