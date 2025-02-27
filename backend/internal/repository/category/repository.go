package category

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	def "github.com/Fi44er/sdmedik/backend/internal/repository"
	"github.com/Fi44er/sdmedik/backend/pkg/constants"
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

func (r *repository) GetAll(ctx context.Context) (*[]model.Category, error) {
	r.logger.Info("Fetching categories...")
	categories := new([]model.Category)
	if err := r.db.WithContext(ctx).Preload("Characteristics").Preload("Images").Find(categories).Error; err != nil {
		r.logger.Errorf("Failed to fetch categories: %v", err)
		return nil, err
	}
	r.logger.Info("Categories fetched successfully")
	return categories, nil
}

func (r *repository) GetByID(ctx context.Context, id int) (*model.Category, error) {
	r.logger.Info("Fetching category by id...")
	category := new(model.Category)
	if err := r.db.WithContext(ctx).Preload("Characteristics").First(category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		r.logger.Errorf("Failed to fetch category by id: %v", err)
		return nil, err
	}

	r.logger.Info("Category fetched by id successfully")
	return category, nil
}

func (r *repository) GetByName(ctx context.Context, name string) (*model.Category, error) {
	r.logger.Info("Fetching category by name...")
	category := new(model.Category)
	if err := r.db.WithContext(ctx).Find(category, "name = ?", name).Error; err != nil {
		r.logger.Errorf("Failed to fetch category by name: %v", err)
		return nil, err
	}

	if category.ID == 0 {
		return nil, nil
	}

	r.logger.Info("Category fetched by name successfully")
	return category, nil
}

func (r *repository) Delete(ctx context.Context, id int, tx *gorm.DB) error {
	r.logger.Info("Deleting category...")
	db := tx
	if db == nil {
		r.logger.Error("Transaction is nil")
		db = r.db
	}
	result := db.WithContext(ctx).Delete(&model.Category{}, id)
	if err := result.Error; err != nil {
		r.logger.Errorf("Failed to delete category: %v", err)
		return err
	}
	if result.RowsAffected == 0 {
		r.logger.Warnf("Category with ID %v not found", id)
		return constants.ErrCategoryNotFound
	}
	r.logger.Info("Category deleted successfully")
	return nil
}

func (r *repository) GetByIDs(ctx context.Context, ids []int) (*[]model.Category, error) {
	r.logger.Info("Fetching categories by ids...")
	categories := new([]model.Category)
	if len(ids) == 0 {
		return categories, nil
	}
	if err := r.db.WithContext(ctx).Preload("Characteristics").Find(categories, ids).Error; err != nil {
		r.logger.Errorf("Failed to fetch categories by ids: %v", err)
		return nil, err
	}
	r.logger.Info("Categories fetched by ids successfully")
	return categories, nil
}

func (r *repository) Update(ctx context.Context, cateegory *model.Category, tx *gorm.DB) error {
	r.logger.Info("Updating category...")
	db := tx
	if db == nil {
		db = r.db
	}

	if err := db.WithContext(ctx).Model(cateegory).Updates(cateegory).Error; err != nil {
		r.logger.Errorf("Failed to update category: %v", err)
		return err
	}

	r.logger.Info("Category updated successfully")
	return nil
}
