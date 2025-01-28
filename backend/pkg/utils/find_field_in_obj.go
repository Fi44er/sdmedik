package utils

import (
	"fmt"
	"reflect"
)

func FindFieldInObject(obj interface{}, fieldName string) (interface{}, error) {
	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("объект не является структурой или указателем на структуру")
	}

	field := val.FieldByName(fieldName)
	if !field.IsValid() {
		return nil, fmt.Errorf("поле '%s' не найдено в объекте", fieldName)
	}

	return field.Interface(), nil
}
