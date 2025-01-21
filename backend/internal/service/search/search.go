package search

import (
	"context"
	"strings"

	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/blevesearch/bleve/v2"
)

const indexDir = "product_index" // Директория для хранения индекса

func (s *service) Search(ctx context.Context, query string) (*[]response.SearchRes, error) {
	words := strings.Fields(query) // Разбиваем запрос на слова
	var resp []response.SearchRes

	// Создаем BooleanQuery
	booleanQuery := bleve.NewBooleanQuery()
	for _, word := range words {
		// prefixQuery := bleve.NewPrefixQuery(word)
		// prefixQuery.SetField("Name")      // Ищем по полю "Name"
		// booleanQuery.AddMust(prefixQuery) // Все префиксы должны совпадать

		fuzzyQuery := bleve.NewFuzzyQuery(word)
		fuzzyQuery.SetField("Name")
		booleanQuery.AddShould(fuzzyQuery)
	}

	searchRequest := bleve.NewSearchRequest(booleanQuery)
	searchRequest.Fields = []string{"Name", "Type"} // Указываем, какие поля включить в результаты
	index := s.indexService.Get()
	searchResult, err := index.Search(searchRequest)
	if err != nil {
		s.logger.Fatalf("Ошибка при поиске: %v", err)
		return nil, err
	}

	// Выводим результаты поиска
	for _, hit := range searchResult.Hits {
		name := hit.Fields["Name"].(string)
		typeElm := hit.Fields["Type"].(string)

		element := response.SearchRes{
			ID:   hit.ID,
			Name: name,
			Type: typeElm,
		}

		if typeElm == "category" {
			resp = append([]response.SearchRes{element}, resp...)
		} else {
			// Иначе добавляем в конец среза
			resp = append(resp, element)
		}
	}

	return &resp, nil
}
