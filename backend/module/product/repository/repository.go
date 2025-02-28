package repository

import (
	"context"

	file_repository "github.com/Fi44er/sdmedik/backend/module/file/repository"
	"github.com/Fi44er/sdmedik/backend/module/product/domain"
	"github.com/Fi44er/sdmedik/backend/shared/logger"
	"gorm.io/gorm"
)

var _ IProductRepository = (*ProductRepository)(nil)

type IProductRepository interface {
	Create(ctx context.Context, productDomain *domain.Product, tx *gorm.DB) error

	GetByID(ctx context.Context, id string) (*domain.Product, error)
}

type ProductRepository struct {
	logger   *logger.Logger
	db       *gorm.DB
	fileRepo file_repository.IFileRepository
}

func NewProductRepository(
	logger *logger.Logger,
	db *gorm.DB,
	fileRepo file_repository.IFileRepository,
) *ProductRepository {
	return &ProductRepository{
		logger:   logger,
		db:       db,
		fileRepo: fileRepo,
	}
}
