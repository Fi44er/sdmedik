package blog

import (
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

var _ def.IBlogService = (*service)(nil)

type service struct {
	logger *logger.Logger
	db     *gorm.DB
	repo   repository.IBlogRepository
}

func NewService(
	logger *logger.Logger,
	db *gorm.DB,
	repo repository.IBlogRepository,
) *service {
	return &service{
		logger: logger,
		db:     db,
		repo:   repo,
	}
}
