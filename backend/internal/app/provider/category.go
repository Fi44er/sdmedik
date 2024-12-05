package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/api/category"
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	categoryRepository "github.com/Fi44er/sdmedik/backend/internal/repository/category"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	categoryService "github.com/Fi44er/sdmedik/backend/internal/service/category"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CategoryProvider struct {
	categoryRepository repository.ICategoryRepository
	categoryService    service.ICategoryService
	categoryImpl       *category.Implementation

	logger    *logger.Logger
	db        *gorm.DB
	validator *validator.Validate

	characteristicService service.ICharacteristicService
}

func NewCategoryProvider(
	logger *logger.Logger,
	db *gorm.DB,
	validator *validator.Validate,
	characteristicService service.ICharacteristicService,
) *CategoryProvider {
	return &CategoryProvider{
		logger:                logger,
		db:                    db,
		validator:             validator,
		characteristicService: characteristicService,
	}
}

func (p *CategoryProvider) CategoryRepository() repository.ICategoryRepository {
	if p.categoryRepository == nil {
		p.categoryRepository = categoryRepository.NewRepository(p.logger, p.db)
	}
	return p.categoryRepository
}

func (p *CategoryProvider) CategoryService() service.ICategoryService {
	if p.categoryService == nil {
		p.categoryService = categoryService.NewService(p.CategoryRepository(), p.logger, p.validator, p.characteristicService)
	}

	return p.categoryService
}

func (p *CategoryProvider) CategoryImpl() *category.Implementation {
	if p.categoryImpl == nil {
		p.categoryImpl = category.NewImplementation(p.CategoryService())
	}

	return p.categoryImpl
}
