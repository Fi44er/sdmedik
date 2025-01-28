package utils

import (
	"fmt"
	"reflect"
)

// Convert преобразует объект src в объект dest.
// Оба объекта должны быть указателями на структуры.
func DtoToModel(src interface{}, dest interface{}) error {
	srcVal := reflect.ValueOf(src).Elem()
	destVal := reflect.ValueOf(dest).Elem()

	if srcVal.Kind() != reflect.Struct || destVal.Kind() != reflect.Struct {
		return fmt.Errorf("src and dest must be structs")
	}

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Type().Field(i)
		destField := destVal.FieldByName(srcField.Name)

		// Если поле в dest существует и может быть установлено
		if destField.IsValid() && destField.CanSet() {
			srcFieldValue := srcVal.Field(i)

			// Проверяем, совместимы ли типы
			if srcFieldValue.Type().AssignableTo(destField.Type()) {
				destField.Set(srcFieldValue)
			} else {
				// Если типы не совместимы, пытаемся преобразовать
				if srcFieldValue.Type().ConvertibleTo(destField.Type()) {
					destField.Set(srcFieldValue.Convert(destField.Type()))
				} else {
					// Пропускаем поле, если типы не совместимы и не конвертируемы
					continue
				}
			}
		}
	}

	return nil
}
