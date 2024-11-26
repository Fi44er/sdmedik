package product

import (
	"context"

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

func NewRepository(logger *logger.Logger, db *gorm.DB) *repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (s *repository) Create(ctx context.Context, data *model.Product) error {
	s.logger.Info("Creating product...")
	if err := s.db.WithContext(ctx).Create(data).Error; err != nil {
		s.logger.Errorf("Failed to create product: %v", err)
		return err
	}

	s.logger.Infof("Product created successfully")
	return nil
}

func (s *repository) GetByID(ctx context.Context, id string) (model.Product, error) {
	s.logger.Infof("Fetching product with ID: %s...", id)
	var product model.Product
	if err := s.db.WithContext(ctx).Where("id = ?", id).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Warnf("Product with ID %s not found", id)
			return product, nil
		}
		s.logger.Errorf("Failed to fetch product with ID %s: %v", id, err)
		return model.Product{}, err
	}
	s.logger.Info("Product fetched successfully")
	return product, nil
}

func (s *repository) GetAll(ctx context.Context, offset int, limit int) ([]model.Product, error) {
	s.logger.Info("Fetching products...")
	var products []model.Product
	if offset == 0 {
		offset = -1
	}

	if limit == 0 {
		limit = -1
	}

	if err := s.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		s.logger.Errorf("Failed to fetch products: %v", err)
		return nil, err
	}
	s.logger.Info("Products fetched successfully")
	return products, nil
}

func (s *repository) Update(ctx context.Context, data *model.Product) error {
	s.logger.Info("Updating product...")
	if err := s.db.WithContext(ctx).Model(data).Updates(data).Error; err != nil {
		s.logger.Errorf("Failed to update product: %v", err)
		return err
	}
	s.logger.Info("Product updated successfully")
	return nil
}

func (s *repository) Delete(ctx context.Context, id string) error {
	s.logger.Infof("Deleting product with ID: %s...", id)
	if err := s.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Product{}).Error; err != nil {
		s.logger.Errorf("Failed to delete product: %v", err)
		return err
	}
	s.logger.Infof("Product deleted by ID: %v successfully", id)
	return nil
}
