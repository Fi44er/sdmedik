package search

import (
	"context"

	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/blevesearch/bleve/v2"
	"github.com/go-playground/validator/v10"
)

var _ def.ISearchService = (*service)(nil)

type service struct {
	logger    *logger.Logger
	validator *validator.Validate

	productService  def.IProductService
	categoryService def.ICategoryService

	index bleve.Index
}

func NewService(
	logger *logger.Logger,
	validator *validator.Validate,
	productService def.IProductService,
	categoryService def.ICategoryService,
) (*service, error) { // Возвращаем ошибку, если инициализация не удалась
	// Создаем или загружаем индекс
	index, err := createOrLoadIndex()
	if err != nil {
		logger.Fatalf("Ошибка при создании/загрузке индекса: %v", err)
		return nil, err
	}

	// Создаем сервис
	svc := &service{
		logger:          logger,
		validator:       validator,
		productService:  productService,
		categoryService: categoryService,

		index: index,
	}

	// Добавляем товары в индекс
	if err := svc.addSampleProducts(context.Background(), index); err != nil {
		logger.Fatalf("Ошибка при добавлении товаров: %v", err)
		return nil, err
	}

	return svc, nil
}
