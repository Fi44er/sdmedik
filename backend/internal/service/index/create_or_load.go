package index

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/constants"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
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

	if products == nil {
		products = &[]response.ProductResponse{}
	}

	categories, err := s.categoryService.GetAll(ctx)
	if err != nil {
		if !errors.Is(err, constants.ErrCategoryNotFound) {
			return err
		}
	}

	if categories == nil {
		categories = &[]model.Category{}
	}

	// Создаем пакет для индексации
	batch := s.index.NewBatch()

	// Добавляем товары в пакет
	for _, product := range *products {
		if err := s.addToBatch(batch, product, "product"); err != nil {
			s.logger.Errorf("ошибка при добавлении товара с ID %s в пакет: %v", product.ID, err)
		}
	}

	// Добавляем категории в пакет
	for _, category := range *categories {
		if err := s.addToBatch(batch, category, "product"); err != nil {
			s.logger.Errorf("ошибка при добавлении категории с ID %v в пакет: %v", category.ID, err)
		}
	}

	// Выполняем пакетную индексацию
	if err := s.index.Batch(batch); err != nil {
		return fmt.Errorf("ошибка при выполнении пакетной индексации: %v", err)
	}

	return nil
}

func (s *service) addToBatch(batch *bleve.Batch, data interface{}, docType string) error {
	name, err := utils.FindFieldInObject(data, "Name")
	if err != nil {
		return err
	}

	id, err := utils.FindFieldInObject(data, "ID")
	if err != nil {
		return err
	}

	strId, err := utils.StringifyID(id)
	if err != nil {
		return err
	}

	doc := map[string]interface{}{
		"Name": name,
		"Type": docType,
	}

	batch.Index(strId, doc)
	return nil
}
