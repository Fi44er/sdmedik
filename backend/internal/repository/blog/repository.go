package blog

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	def "github.com/Fi44er/sdmedik/backend/internal/repository"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

var _ def.IBlogRepository = (*repository)(nil)

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

func (r *repository) Create(ctx context.Context, data *model.Blog) error {
	r.logger.Info("Creating blog...")

	if err := r.db.WithContext(ctx).Create(data).Error; err != nil {
		r.logger.Errorf("Failed to create blog: %v", err)
		return err
	}
	r.logger.Info("Blog creating successfully")
	return nil
}

func (r *repository) GetAll(ctx context.Context, offset, limit int) ([]model.Blog, error) {
	r.logger.Info("Fetching blog...")
	if offset <= 0 {
		offset = -1
	}
	if limit <= 0 {
		limit = -1 // или другое значение по умолчанию
	}
	blogs := new([]model.Blog)
	if err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(blogs).Error; err != nil {
		r.logger.Errorf("Failed to fetch blogs: %v", err)
		return nil, err
	}

	r.logger.Info("Blogs fetching successfully")

	return *blogs, nil
}

func (r *repository) GetByID(ctx context.Context, id string) (*model.Blog, error) {
	r.logger.Info("Fetching blog")
	blog := new(model.Blog)
	if err := r.db.WithContext(ctx).First(blog, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		r.logger.Errorf("Failed to fetch blog by id: %v", err)
		return nil, err
	}

	r.logger.Info("Blog fetched by id successfully")
	return blog, nil
}

func (r *repository) Update(ctx context.Context, data *model.Blog) error {
	r.logger.Info("Updating blog")
	if err := r.db.WithContext(ctx).Model(data).Updates(data).Error; err != nil {
		r.logger.Errorf("Failed to update blog: %v", err)
		return err
	}

	r.logger.Info("Blog updated successfully")
	return nil
}
