package repository

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/module/category/domain"
	file_repository "github.com/Fi44er/sdmedik/backend/module/file/repository"
	"github.com/Fi44er/sdmedik/backend/shared/logger"
	"gorm.io/gorm"
)

var _ ICategoryRepository = (*CategoryRepository)(nil)

type ICategoryRepository interface {
	Create(ctx context.Context, categoryDomain *domain.Category, tx *gorm.DB) error

	GetByID(ctx context.Context, id string) (*domain.Category, error)
	GetAll(ctx context.Context) ([]domain.Category, error)
	GetByIDs(ctx context.Context, ids []string) ([]domain.Category, error)
}

type CategoryRepository struct {
	logger   *logger.Logger
	db       *gorm.DB
	fileRepo file_repository.IFileRepository
}

func NewCategoryRepository(
	logger *logger.Logger,
	db *gorm.DB,
	fileRepo file_repository.IFileRepository,
) *CategoryRepository {
	return &CategoryRepository{
		logger:   logger,
		db:       db,
		fileRepo: fileRepo,
	}
}
