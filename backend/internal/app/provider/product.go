package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/api/product"
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	productRepository "github.com/Fi44er/sdmedik/backend/internal/repository/product"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	productService "github.com/Fi44er/sdmedik/backend/internal/service/product"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProductProvider struct {
	productRepository repository.IProductRepository
	productService    service.IProductService
	productImpl       *product.Implementation

	logger    *logger.Logger
	db        *gorm.DB
	validator *validator.Validate

	categoryService service.ICategoryService
}

func NewProductProvider(
	logger *logger.Logger,
	db *gorm.DB,
	validator *validator.Validate,
	categoryService service.ICategoryService,
) *ProductProvider {
	return &ProductProvider{
		logger:          logger,
		db:              db,
		validator:       validator,
		categoryService: categoryService,
	}
}

func (p *ProductProvider) ProductRepository() repository.IProductRepository {
	if p.productRepository == nil {
		p.productRepository = productRepository.NewRepository(p.logger, p.db)
	}
	return p.productRepository
}

func (p *ProductProvider) ProductService() service.IProductService {
	if p.productService == nil {
		p.productService = productService.NewService(p.ProductRepository(), p.logger, p.validator, p.categoryService)
	}

	return p.productService
}

func (p *ProductProvider) ProductImpl() *product.Implementation {
	if p.productImpl == nil {
		p.productImpl = product.NewImplementation(p.ProductService())
	}

	return p.productImpl
}