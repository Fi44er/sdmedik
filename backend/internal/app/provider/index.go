package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/service"
	indexService "github.com/Fi44er/sdmedik/backend/internal/service/index"
	events "github.com/Fi44er/sdmedik/backend/pkg/evenbus"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
)

type IndexProvider struct {
	indexService service.IIndexService

	logger    *logger.Logger
	validator *validator.Validate
	evetBus   *events.EventBus

	productService  service.IProductService
	categoryService service.ICategoryService
}

func NewIndexProvider(
	logger *logger.Logger,
	validator *validator.Validate,
	productService service.IProductService,
	categoryService service.ICategoryService,
	eventBus *events.EventBus,
) *IndexProvider {
	return &IndexProvider{
		logger:          logger,
		validator:       validator,
		productService:  productService,
		categoryService: categoryService,
		evetBus:         eventBus,
	}
}

func (p *IndexProvider) IndexService() service.IIndexService {
	if p.indexService == nil {
		var err error
		p.indexService, err = indexService.NewService(p.logger, p.validator, p.productService, p.categoryService, p.evetBus)
		if err != nil {
			p.logger.Errorf("Error during initializing index service: %s", err.Error())
			return nil
		}
	}
	return p.indexService
}
