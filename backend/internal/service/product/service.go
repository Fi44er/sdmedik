package product

import (
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
)

var _ def.IProductService = (*service)(nil)

type service struct {
	logger    *logger.Logger
	repo      repository.IProductRepository
	validator *validator.Validate

	categoryService def.ICategoryService
}

func NewService(repo repository.IProductRepository, logger *logger.Logger, validator *validator.Validate, categoryService def.ICategoryService) *service {
	return &service{
		repo:            repo,
		logger:          logger,
		validator:       validator,
		categoryService: categoryService,
	}
}
