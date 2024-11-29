package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/api/product"
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	productRepository "github.com/Fi44er/sdmedik/backend/internal/repository/product"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	productService "github.com/Fi44er/sdmedik/backend/internal/service/product"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

type ProductProvider struct {
	productRepository repository.IProductRepository
	productService    service.IProductService
	productImpl       *product.Implementation

	logger *logger.Logger
	db     *gorm.DB
}

func NewProductProvider(logger *logger.Logger, db *gorm.DB) *ProductProvider {
	return &ProductProvider{
		logger: logger,
		db:     db,
	}
}

func (s *ProductProvider) ProductRepository() repository.IProductRepository {
	if s.productRepository == nil {
		s.productRepository = productRepository.NewRepository(s.logger, s.db)
	}
	return s.productRepository
}

func (s *ProductProvider) ProductService() service.IProductService {
	if s.productService == nil {
		s.productService = productService.NewService(s.ProductRepository(), s.logger)
	}

	return s.productService
}

func (s *ProductProvider) ProductImpl() *product.Implementation {
	if s.productImpl == nil {
		s.productImpl = product.NewImplementation(s.ProductService())
	}

	return s.productImpl
}
