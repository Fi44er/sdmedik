package index

import (
	"context"

	def "github.com/Fi44er/sdmedik/backend/internal/service"
	events "github.com/Fi44er/sdmedik/backend/pkg/evenbus"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/blevesearch/bleve/v2"
	"github.com/go-playground/validator/v10"
)

var _ def.IIndexService = (*service)(nil)

type service struct {
	logger    *logger.Logger
	validator *validator.Validate
	eventBus  *events.EventBus

	index bleve.Index

	productService  def.IProductService
	categoryService def.ICategoryService
}

func NewService(
	logger *logger.Logger,
	validator *validator.Validate,
	productService def.IProductService,
	categoryService def.ICategoryService,
	eventBus *events.EventBus,
) (*service, error) {
	svc := &service{
		logger:          logger,
		validator:       validator,
		productService:  productService,
		categoryService: categoryService,
		eventBus:        eventBus,
	}

	index, err := svc.CreateOrLoad()
	if err != nil {
		logger.Fatalf("Ошибка при создании/загрузке индекса: %v", err)
		return nil, err
	}

	svc.index = index
	eventBus.Subscribe(events.EventDataCreatedOrUpdated, svc.handleDataCreatedOrUpdated)
	eventBus.Subscribe(events.EventDataDeleted, svc.handleDataDeleted)

	if err := svc.addSampleProducts(context.Background()); err != nil {
		logger.Fatalf("Ошибка при добавлении товаров: %v", err)
		return nil, err
	}

	return svc, nil
}

func (s *service) handleDataCreatedOrUpdated(event events.Event) {
	data := event.Data
	dataType := event.DataType
	if err := s.AddOrUpdate(data, dataType); err != nil {
		s.logger.Errorf("Ошибка при индексации категории: %v", err)
	}
}

func (s *service) handleDataDeleted(event events.Event) {
	data := event.Data
	if err := s.Delete(data); err != nil {
		s.logger.Errorf("Ошибка при удалении товара: %v", err)
	}
}
