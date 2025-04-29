package utils

import (
	"log"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func ValidateCharacteristicValue(categories []model.Category, values []dto.CharacteristicValue) error {
	characteristicsMap := make(map[int]string)
	for _, category := range categories {
		for _, characteristic := range category.Characteristics {
			characteristicsMap[characteristic.ID] = string(characteristic.DataType)
		}
	}

	for _, value := range values {
		_, _ = characteristicsMap[value.CharacteristicID]
		log.Printf("%+v", categories)
		// switch characteristicType {
		// case "string":
		// 	continue
		// case "int":
		// 	_, err := strconv.Atoi(value.Value)
		// 	if err != nil {
		// 		return errors.New(400, "invalid value for int type")
		// 	}
		// case "float":
		// 	_, err := strconv.ParseFloat(value.Value, 64)
		// 	if err != nil {
		// 		return errors.New(400, "invalid value for float type")
		// 	}
		// case "bool":
		// 	_, err := strconv.ParseBool(value.Value)
		// 	if err != nil {
		// 		return errors.New(400, "invalid value for bool type")
		// 	}
		// default:
		// 	return errors.New(400, "unsupported data type")
		// }
	}

	return nil
}
