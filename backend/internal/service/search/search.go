package search

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/custom"
	"github.com/blevesearch/bleve/v2/analysis/token/lowercase"
	"github.com/blevesearch/bleve/v2/analysis/token/ngram"
	"github.com/blevesearch/bleve/v2/mapping"
)

const indexDir = "product_index" // Директория для хранения индекса

func (s *service) Search(ctx context.Context, query string) (*[]response.SearchRes, error) {
	words := strings.Fields(query) // Разбиваем запрос на слова
	var resp []response.SearchRes

	// Создаем BooleanQuery
	booleanQuery := bleve.NewBooleanQuery()
	for _, word := range words {
		matchQuery := bleve.NewMatchQuery(word)
		booleanQuery.AddMust(matchQuery)
	}

	searchRequest := bleve.NewSearchRequest(booleanQuery)
	searchRequest.Fields = []string{"Name", "Price"} // Указываем, какие поля включить в результаты
	searchResult, err := s.index.Search(searchRequest)
	if err != nil {
		s.logger.Fatalf("Ошибка при поиске: %v", err)
		return nil, err
	}

	// Выводим результаты поиска
	for _, hit := range searchResult.Hits {
		name := hit.Fields["Name"]
		resp = append(resp, response.SearchRes{ID: hit.ID, Name: name.(string)})
	}

	return &resp, nil
}

func createOrLoadIndex() (bleve.Index, error) {
	if _, err := os.Stat(indexDir); os.IsNotExist(err) {
		return createIndex()
	}

	return bleve.Open(indexDir)
}

func createIndex() (bleve.Index, error) {
	indexMapping := bleve.NewIndexMapping()
	err := indexMapping.AddCustomTokenFilter("ngram3", map[string]interface{}{
		"type": ngram.Name,
		"min":  2,
		"max":  3,
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка при регистрации токенизатора: %v", err)
	}

	err = indexMapping.AddCustomTokenFilter("lowercase", map[string]interface{}{
		"type": lowercase.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка при регистрации lowercase фильтра: %v", err)
	}

	err = indexMapping.AddCustomAnalyzer("custom", map[string]interface{}{
		"type":      custom.Name,
		"tokenizer": "unicode", // Используем стандартный токенизатор
		"token_filters": []string{
			"lowercase", // Приводим текст к нижнему регистру
			"ngram3",    // Применяем n-gram фильтр
		},
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка при добавлении кастомного анализатора: %v", err)
	}
	indexMapping.DefaultAnalyzer = "custom"

	productMapping := bleve.NewDocumentMapping()
	productMapping.AddFieldMappingsAt("Name", mapping.NewTextFieldMapping())

	indexMapping.AddDocumentMapping("product", productMapping)

	index, err := bleve.New(indexDir, indexMapping)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании индекса: %v", err)
	}

	return index, nil
}

func (s *service) addSampleProducts(ctx context.Context, index bleve.Index) error {
	// Генерируем товары
	products, err := s.productService.Get(ctx, dto.ProductSearchCriteria{Minimal: true})
	if err != nil {
		return err
	}
	// Добавляем только те товары, которых нет в индексе
	for _, product := range *products {
		// Проверяем, существует ли товар с таким ID в индексе
		doc, err := index.Document(product.ID)
		if err != nil {
			return fmt.Errorf("ошибка при проверке документа с ID %s: %v", product.ID, err)
		}

		// Если товара нет в индексе, добавляем его
		if doc == nil {
			if err := index.Index(product.ID, map[string]interface{}{
				"Name": product.Name,
			}); err != nil {
				return fmt.Errorf("ошибка при индексации товара с ID %s: %v", product.ID, err)
			}
		}
	}

	return nil
}
