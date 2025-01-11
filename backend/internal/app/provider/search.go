package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/api/search"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	searchService "github.com/Fi44er/sdmedik/backend/internal/service/search"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
)

type SearchProvider struct {
	searchService service.ISearchService
	searchImpl    *search.Implementation

	logger    *logger.Logger
	validator *validator.Validate

	productService  service.IProductService
	categoryService service.ICategoryService
}

func NewSearchProvider(
	logger *logger.Logger,
	validator *validator.Validate,
	productService service.IProductService,
	categoryService service.ICategoryService,
) *SearchProvider {
	return &SearchProvider{
		logger:          logger,
		validator:       validator,
		productService:  productService,
		categoryService: categoryService,
	}
}

func (p *SearchProvider) SearchService() service.ISearchService {
	if p.searchService == nil {
		var err error
		p.searchService, err = searchService.NewService(p.logger, p.validator, p.productService, p.categoryService)
		if err != nil {
			p.logger.Errorf("Error during initializing search service: %s", err.Error())
			return nil
		}
	}
	return p.searchService
}

func (p *SearchProvider) SearchImpl() *search.Implementation {
	if p.searchImpl == nil {
		p.searchImpl = search.NewImplementation(p.SearchService())
	}
	return p.searchImpl
}
