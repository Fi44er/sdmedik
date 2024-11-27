package utils

import (
	"reflect"
)

// Convert преобразует объект src в объект dest.
// Оба объекта должны быть указателями на структуры.
func DtoToModel(src interface{}, dest interface{}) error {
	srcVal := reflect.ValueOf(src).Elem()
	destVal := reflect.ValueOf(dest).Elem()

	if srcVal.Kind() != reflect.Struct || destVal.Kind() != reflect.Struct {
		return nil // или верните ошибку, если нужно
	}

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Type().Field(i)
		destField := destVal.FieldByName(srcField.Name)

		if destField.IsValid() && destField.CanSet() {
			destField.Set(srcVal.Field(i))
		}
	}

	return nil
}
