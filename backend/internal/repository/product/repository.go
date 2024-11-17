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
	return s.db.WithContext(ctx).Create(data).Error
}

func (s *repository) GetByID(ctx context.Context, id string) (model.Product, error) {
	var product model.Product
	if err := s.db.WithContext(ctx).Where("id = ?", id).First(&product).Error; err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (s *repository) GetAll(ctx context.Context, offset int, limit int) ([]model.Product, error) {
	var products []model.Product
	switch {

	case offset == 0 && limit == 0:
		offset = -1
		limit = -1
	case offset == 0:
		offset = -1
	case limit == 0:
		limit = -1
	default:
		offset *= limit
	}

	if err := s.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (s *repository) Update(ctx context.Context, data *model.Product) error {
	return s.db.WithContext(ctx).Save(data).Error
}

func (s *repository) Delete(ctx context.Context, id string) error {
	return s.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Product{}).Error
}
