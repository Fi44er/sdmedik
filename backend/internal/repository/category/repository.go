package category

import (
	"context"
	"fmt"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	def "github.com/Fi44er/sdmedik/backend/internal/repository"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

var _ def.ICategoryRepository = (*repository)(nil)

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

func (r *repository) Create(ctx context.Context, data *model.Category, tx *gorm.DB) error {
	r.logger.Info("Creating category...")
	db := tx
	if db == nil {
		db = r.db
	}

	if err := db.WithContext(ctx).Create(data).Error; err != nil {
		r.logger.Errorf("Failed to create category: %v", err)
		return err
	}

	r.logger.Infof("Category created successfully")
	return nil
}

func (r *repository) GetAll(ctx context.Context) ([]model.Category, error) {

	r.logger.Info("Fetching categories...")
	var categories []model.Category
	if err := r.db.WithContext(ctx).Preload("Characteristics").Find(&categories).Error; err != nil {
		r.logger.Errorf("Failed to fetch categories: %v", err)
		return nil, err
	}
	r.logger.Info("Categories fetched successfully")
	return categories, nil
}

func (r *repository) GetByID(ctx context.Context, id int) (model.Category, error) {
	r.logger.Info("Fetching category...")
	var category model.Category
	if err := r.db.WithContext(ctx).Preload("Products").First(&category, id).Error; err != nil {
		r.logger.Errorf("Failed to fetch category: %v", err)
		return model.Category{}, err
	}
	r.logger.Info("Category fetched successfully")
	return category, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	r.logger.Info("Deleting category...")
	result := r.db.WithContext(ctx).Delete(&model.Category{}, id)
	if err := result.Error; err != nil {
		r.logger.Errorf("Failed to delete category: %v", err)
		return err
	}

	if result.RowsAffected == 0 {
		r.logger.Warnf("Category with ID %v not found", id)
		return fmt.Errorf("Category not found")
	}

	r.logger.Info("Category deleted successfully")
	return nil
}

func (r *repository) GetByIDs(ctx context.Context, ids []int) ([]model.Category, error) {
	r.logger.Info("Fetching categories...")
	var categories []model.Category
	if err := r.db.WithContext(ctx).Preload("Characteristics").Find(&categories, ids).Error; err != nil {
		r.logger.Errorf("Failed to fetch categories: %v", err)
		return nil, err
	}
	r.logger.Info("Categories fetched successfully")
	return categories, nil
}

func (r *repository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}
