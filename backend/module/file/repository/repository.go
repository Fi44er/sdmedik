package repository

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/module/file/domain"
	"github.com/Fi44er/sdmedik/backend/shared/logger"
	"gorm.io/gorm"
)

var _ IFileRepository = (*FileRepository)(nil)

type IFileRepository interface {
	CreateMany(ctx context.Context, filesDomains []domain.File, tx *gorm.DB) error
	GetFilesByOwner(ctx context.Context, ownerID string, ownerType string) ([]domain.File, error)
}

type FileRepository struct {
	logger *logger.Logger
	db     *gorm.DB
}

func NewFileRepository(
	logger *logger.Logger,
	db *gorm.DB,
) *FileRepository {
	return &FileRepository{
		logger: logger,
		db:     db,
	}
}
