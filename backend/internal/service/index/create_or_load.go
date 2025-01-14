package index

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/constants"
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/custom"
	"github.com/blevesearch/bleve/v2/analysis/token/lowercase"
	"github.com/blevesearch/bleve/v2/analysis/tokenizer/whitespace"
	"github.com/blevesearch/bleve/v2/mapping"
)

const indexDir = "product_index" // Директория для хранения индекса

func (s *service) CreateOrLoad() (bleve.Index, error) {
	if _, err := os.Stat(indexDir); os.IsNotExist(err) {
		return s.createIndex()
	}

	return bleve.Open(indexDir)
}

func (s *service) createIndex() (bleve.Index, error) {
	indexMapping := bleve.NewIndexMapping()

	var err error

	err = indexMapping.AddCustomTokenizer("whitespace", map[string]interface{}{
		"type": whitespace.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка при регистрации токенизатора whitespace: %v", err)
	}

	err = indexMapping.AddCustomTokenFilter("lowercase", map[string]interface{}{
		"type": lowercase.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка при регистрации lowercase фильтра: %v", err)
	}
	// Добавляем кастомный анализатор
	err = indexMapping.AddCustomAnalyzer("prefix_analyzer", map[string]interface{}{
		"char_filters": []interface{}{},
		"tokenizer":    `whitespace`,
		"type":         custom.Name,
		"token_filters": []string{
			"lowercase", // Приводим текст к нижнему регистру
		},
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка при добавлении кастомного анализатора: %v", err)
	}
	indexMapping.DefaultAnalyzer = "prefix_analyzer"

	// Создаем маппинг для продукта
	productMapping := bleve.NewDocumentMapping()
	nameFieldMapping := mapping.NewTextFieldMapping()
	nameFieldMapping.Analyzer = "prefix_analyzer" // Используем кастомный анализатор для поля "Name"
	productMapping.AddFieldMappingsAt("Name", nameFieldMapping)
	productMapping.AddFieldMappingsAt("Type", mapping.NewTextFieldMapping())

	indexMapping.AddDocumentMapping("product", productMapping)

	// Создаем индекс
	index, err := bleve.New(indexDir, indexMapping)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании индекса: %v", err)
	}

	return index, nil
}

func (s *service) addSampleProducts(ctx context.Context) error {
	// Генерируем товары
	products, err := s.productService.Get(ctx, dto.ProductSearchCriteria{Minimal: true})
	if err != nil {
		return err
	}

	categories, err := s.categoryService.GetAll(ctx)
	if err != nil {
		if !errors.Is(err, constants.ErrCategoryNotFound) {
			return err
		}
	}

	for _, product := range *products {
		if err := s.AddOrUpdate(product, "product"); err != nil {
			s.logger.Errorf("ошибка при индексации товара с ID %s: %v", product.ID, err)
		}
	}

	for _, category := range *categories {
		if err := s.AddOrUpdate(category, "product"); err != nil {
			s.logger.Errorf("ошибка при индексации товара с ID %v: %v", category.ID, err)
		}
	}

	return nil
}
