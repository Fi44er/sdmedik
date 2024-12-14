package utils

import (
	"strconv"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func ValidateCharacteristicValue(categories []model.Category, values []dto.CharacteristicValue) error {
	characteristicsMap := make(map[int]string)
	for _, category := range categories {
		for _, characteristic := range category.Characteristics {
			characteristicsMap[characteristic.ID] = string(characteristic.DataType)
		}
	}

	for _, value := range values {
		characteristicType, _ := characteristicsMap[value.CharacteristicID]
		switch characteristicType {
		case "string":
			continue
		case "int":
			_, err := strconv.Atoi(value.Value)
			if err != nil {
				return errors.New(400, "invalid value for int type")
			}
		case "float":
			_, err := strconv.ParseFloat(value.Value, 64)
			if err != nil {
				return errors.New(400, "invalid value for float type")
			}
		case "bool":
			_, err := strconv.ParseBool(value.Value)
			if err != nil {
				return errors.New(400, "invalid value for bool type")
			}
		default:
			return errors.New(400, "unsupported data type")
		}
	}

	return nil
}
