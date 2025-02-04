package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/api/category"
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	categoryRepository "github.com/Fi44er/sdmedik/backend/internal/repository/category"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	categoryService "github.com/Fi44er/sdmedik/backend/internal/service/category"
	events "github.com/Fi44er/sdmedik/backend/pkg/evenbus"
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
	eventBus  *events.EventBus

	characteristicService  service.ICharacteristicService
	transactionManagerRepo repository.ITransactionManager
	imageService           service.IImageService
}

func NewCategoryProvider(
	logger *logger.Logger,
	db *gorm.DB,
	validator *validator.Validate,
	characteristicService service.ICharacteristicService,
	transactionManagerRepo repository.ITransactionManager,
	imageService service.IImageService,
	eventBus *events.EventBus,
) *CategoryProvider {
	return &CategoryProvider{
		logger:                 logger,
		db:                     db,
		validator:              validator,
		characteristicService:  characteristicService,
		transactionManagerRepo: transactionManagerRepo,
		imageService:           imageService,
		eventBus:               eventBus,
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
		p.categoryService = categoryService.NewService(
			p.CategoryRepository(),
			p.logger,
			p.validator,
			p.characteristicService,
			p.transactionManagerRepo,
			p.imageService,
			p.eventBus,
		)
	}

	return p.categoryService
}

func (p *CategoryProvider) CategoryImpl() *category.Implementation {
	if p.categoryImpl == nil {
		p.categoryImpl = category.NewImplementation(p.CategoryService())
	}

	return p.categoryImpl
}
