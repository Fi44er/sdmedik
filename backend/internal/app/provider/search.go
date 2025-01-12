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

	indexService service.IIndexService
}

func NewSearchProvider(
	logger *logger.Logger,
	validator *validator.Validate,
	indexService service.IIndexService,
) *SearchProvider {
	return &SearchProvider{
		logger:       logger,
		validator:    validator,
		indexService: indexService,
	}
}

func (p *SearchProvider) SearchService() service.ISearchService {
	if p.searchService == nil {
		p.searchService = searchService.NewService(p.logger, p.validator, p.indexService)
	}
	return p.searchService
}

func (p *SearchProvider) SearchImpl() *search.Implementation {
	if p.searchImpl == nil {
		p.searchImpl = search.NewImplementation(p.SearchService())
	}
	return p.searchImpl
}
