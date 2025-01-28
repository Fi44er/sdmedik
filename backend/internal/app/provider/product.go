package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/api/product"
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	productRepository "github.com/Fi44er/sdmedik/backend/internal/repository/product"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	productService "github.com/Fi44er/sdmedik/backend/internal/service/product"
	events "github.com/Fi44er/sdmedik/backend/pkg/evenbus"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ProductProvider struct {
	productRepository repository.IProductRepository
	productService    service.IProductService
	productImpl       *product.Implementation

	logger    *logger.Logger
	db        *gorm.DB
	validator *validator.Validate
	eventBus  *events.EventBus
	cache     *redis.Client

	categoryService            service.ICategoryService
	characteristicValueService service.ICharacteristicValueService
	transactionManagerRepo     repository.ITransactionManager
	imageService               service.IImageService
	characteristicService      service.ICharacteristicService
	certificateService         service.ICertificateService
}

func NewProductProvider(
	logger *logger.Logger,
	db *gorm.DB,
	validator *validator.Validate,
	eventBus *events.EventBus,
	cache *redis.Client,

	categoryService service.ICategoryService,
	characteristicValueService service.ICharacteristicValueService,
	transactionManagerRepo repository.ITransactionManager,
	imageService service.IImageService,
	characteristicService service.ICharacteristicService,
	certificateService service.ICertificateService,
) *ProductProvider {
	return &ProductProvider{
		logger:                     logger,
		db:                         db,
		validator:                  validator,
		categoryService:            categoryService,
		characteristicValueService: characteristicValueService,
		transactionManagerRepo:     transactionManagerRepo,
		imageService:               imageService,
		characteristicService:      characteristicService,
		eventBus:                   eventBus,
		certificateService:         certificateService,
		cache:                      cache,
	}
}

func (p *ProductProvider) ProductRepository() repository.IProductRepository {
	if p.productRepository == nil {
		p.productRepository = productRepository.NewRepository(p.logger, p.db, p.cache)
	}
	return p.productRepository
}

func (p *ProductProvider) ProductService() service.IProductService {
	if p.productService == nil {
		p.productService = productService.NewService(
			p.ProductRepository(),
			p.logger,
			p.validator,
			p.categoryService,
			p.characteristicValueService,
			p.transactionManagerRepo,
			p.imageService,
			p.characteristicService,
			p.eventBus,
			p.certificateService,
		)
	}

	return p.productService
}

func (p *ProductProvider) ProductImpl() *product.Implementation {
	if p.productImpl == nil {
		p.productImpl = product.NewImplementation(p.ProductService())
	}

	return p.productImpl
}
