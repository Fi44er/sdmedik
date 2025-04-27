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

	maxProduct := lo.MaxBy(*products, func(a, b model.Product) bool {
		return a.Price > b.Price
	})

	minProduct := lo.MinBy(*products, func(a, b model.Product) bool {
		return a.Price < b.Price
	})

	filterRes.Price.Max = maxProduct.Price
	filterRes.Price.Min = minProduct.Price

	charMap := make(map[string]struct {
		ID     int
		Type   string
		Values []interface{}
	})

	for _, product := range *products {
		for _, charValue := range product.CharacteristicValues {
			char, err := s.characteristicService.GetByID(ctx, charValue.CharacteristicID)
			if err != nil {
				s.logger.Errorf("Error getting characteristic: %v", err)
				continue
			}

			if char.CategoryID == categoryID {
				var value []interface{}
				// switch char.DataType {
				// case model.TypeInt:
				// 	value, err = strconv.Atoi(charValue.Value)
				// case model.TypeFloat:
				// 	value, err = strconv.ParseFloat(charValue.Value, 64)
				// case model.TypeBool:
				// 	value, err = strconv.ParseBool(charValue.Value)
				// default:
				// 	value = charValue.Value
				// }

				for _, charValue := range charValue.Value {
					if _, exists := charMap[char.Name]; !exists {
						charMap[char.Name] = struct {
							ID     int
							Type   string
							Values []interface{}
						}{
							ID:     char.ID,
							Type:   string(char.DataType),
							Values: []interface{}{},
						}
					}

					if !contains(charMap[char.Name].Values, value) {
						charMap[char.Name] = struct {
							ID     int
							Type   string
							Values []interface{}
						}{
							ID:     char.ID,
							Type:   charMap[char.Name].Type,
							Values: append(charMap[char.Name].Values, charValue),
						}
					}
				}
				// value = charValue.Value

				// if err != nil {
				// 	s.logger.Errorf("Error converting characteristic value: %v", err)
				// 	continue
				// }

			}
		}
	}

	for name, data := range charMap {
		filterRes.Characteristics = append(filterRes.Characteristics, response.CharacteristicFilter{
			ID:     data.ID,
			Name:   name,
			Type:   data.Type,
			Values: data.Values,
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

func contains(slice []interface{}, value interface{}) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
