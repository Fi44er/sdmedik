package repository

import (
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

type FIleRepository struct {
	logger *logger.Logger
	db     *gorm.DB
}

func NewFileRepository(logger *logger.Logger, db *gorm.DB) *FIleRepository {
	return &FIleRepository{
		logger: logger,
		db:     db,
	}
}
