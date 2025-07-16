package page

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	def "github.com/Fi44er/sdmedik/backend/internal/repository"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ def.IPageRepository = (*repository)(nil)

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

func (r *repository) Create(ctx context.Context, data *model.Page) error {
	r.logger.Info("Creating page...")

	if err := r.db.WithContext(ctx).Create(data).Error; err != nil {
		r.logger.Errorf("Failed to create page: %v", err)
		return err
	}

	r.logger.Info("Page created successfully")
	return nil
}

func (r *repository) GetByPath(ctx context.Context, path string) (*model.Page, error) {
	r.logger.Info("Fetching page...")
	page := new(model.Page)
	if err := r.db.WithContext(ctx).Preload("Elements").Where("path = ?", path).First(&page).Error; err != nil {
		r.logger.Errorf("Failed to fetch page: %v", err)
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	r.logger.Info("Page fetched successfully")
	return page, nil
}

func (r *repository) AddOrUpdateElement(ctx context.Context, element *model.Element) error {
	r.logger.Info("Adding or updating element...")

	if err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "element_id"}, {Name: "page_path"}},
		DoUpdates: clause.AssignmentColumns([]string{"value"}),
	}).Create(element).Error; err != nil {
		r.logger.Errorf("Failed to add or update element: %v", err)
		return err
	}

	r.logger.Info("Element added or updated successfully")
	return nil
}
