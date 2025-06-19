package product

import (
	"context"
	"encoding/json"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/samber/lo"
)

func (s *service) GetFilter(ctx context.Context, categoryID int) (*response.ProductFilter, error) {
	filterRes := new(response.ProductFilter)
	_, err := s.categoryService.GetByID(ctx, categoryID)
	if err != nil {
		s.logger.Errorf("Error getting category: %v", err)
		return nil, err
	}

	products, _, err := s.repo.Get(ctx, dto.ProductSearchCriteria{CategoryID: categoryID})
	if err != nil {
		s.logger.Errorf("Error getting products: %v", err)
		return nil, err
	}

	// Получаем мин и макс цену
	maxProduct := lo.MaxBy(*products, func(a, b model.Product) bool {
		return a.Price > b.Price
	})
	minProduct := lo.MinBy(*products, func(a, b model.Product) bool {
		return a.Price < b.Price
	})
	filterRes.Price.Max = maxProduct.Price
	filterRes.Price.Min = minProduct.Price

	// Создаем map для хранения уникальных значений характеристик
	charMap := make(map[string]struct {
		ID     int
		Type   string
		Values map[interface{}]struct{} // Используем map для хранения уникальных значений
	})

	for _, product := range *products {
		for _, charValue := range product.CharacteristicValues {
			char, err := s.characteristicService.GetByID(ctx, charValue.CharacteristicID)
			if err != nil {
				s.logger.Errorf("Error getting characteristic: %v", err)
				continue
			}

			if char.CategoryID == categoryID {
				// Инициализируем структуру для характеристики, если ее еще нет
				if _, exists := charMap[char.Name]; !exists {
					charMap[char.Name] = struct {
						ID     int
						Type   string
						Values map[interface{}]struct{}
					}{
						ID:     char.ID,
						Type:   string(char.DataType),
						Values: make(map[interface{}]struct{}),
					}
				}

				// Добавляем все значения, гарантируя уникальность
				for _, val := range charValue.Value {
					charMap[char.Name].Values[val] = struct{}{}
				}
			}
		}
	}

	// Преобразуем map уникальных значений в slice для ответа
	for name, data := range charMap {
		values := make([]interface{}, 0, len(data.Values))
		for val := range data.Values {
			values = append(values, val)
		}

		filterRes.Characteristics = append(filterRes.Characteristics, response.CharacteristicFilter{
			ID:     data.ID,
			Name:   name,
			Type:   data.Type,
			Values: values,
		})
	}

	filterRes.Count = len(*products)

	jsonData, err := json.MarshalIndent(filterRes, "", "  ")
	if err != nil {
		return nil, err
	}
	s.logger.Infof("Filter: %v", string(jsonData))

	return filterRes, nil
}
