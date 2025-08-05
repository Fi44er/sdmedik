package search

import (
	"context"
	"strings"

	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/blevesearch/bleve/v2"
)

const indexDir = "product_index" // Директория для хранения индекса

func (s *service) Search(ctx context.Context, query string) (*[]response.SearchRes, error) {
	words := strings.Fields(strings.ToLower(query))
	var resp []response.SearchRes

	booleanQuery := bleve.NewBooleanQuery()
	for _, word := range words {
		// Создаем запросы для поиска по названию
		namePrefixQuery := bleve.NewPrefixQuery(word)
		namePrefixQuery.SetField("Name")

		// Создаем запросы для поиска по артикулу
		articlePrefixQuery := bleve.NewTermQuery(word)
		articlePrefixQuery.SetField("Article")

		// Объединяем запросы по названию и артикулу с OR
		nameOrArticleQuery := bleve.NewDisjunctionQuery(namePrefixQuery, articlePrefixQuery)

		// Добавляем в общий запрос
		booleanQuery.AddMust(nameOrArticleQuery)
	}

	searchRequest := bleve.NewSearchRequest(booleanQuery)
	searchRequest.Fields = []string{"Name", "Article", "Type"} // Добавляем Article в возвращаемые поля
	index := s.indexService.Get()
	searchResult, err := index.Search(searchRequest)
	if err != nil {
		s.logger.Fatalf("Ошибка при поиске: %v", err)
		return nil, err
	}

	for _, hit := range searchResult.Hits {
		name := hit.Fields["Name"].(string)
		typeElm := hit.Fields["Type"].(string)
		article := ""
		if art, ok := hit.Fields["Article"].(string); ok {
			article = art
		}

		element := response.SearchRes{
			ID:      hit.ID,
			Name:    name,
			Article: article, // Добавляем артикул в ответ
			Type:    typeElm,
		}

		if typeElm == "category" {
			resp = append([]response.SearchRes{element}, resp...)
		} else {
			resp = append(resp, element)
		}
	}

	return &resp, nil
}
