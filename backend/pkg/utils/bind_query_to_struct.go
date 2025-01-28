package utils

import (
	"encoding/json"
	"log"
	"reflect"
	"strconv"
)

// BindQueryToStruct заполняет структуру dest параметрами из queryParams,
// преобразуя их к соответствующим типам структуры.
func BindQueryToStruct(queryParams map[string]string, dest interface{}) {
	destValue := reflect.ValueOf(dest).Elem()
	destType := reflect.TypeOf(dest).Elem()

	for i := 0; i < destValue.NumField(); i++ {
		field := destType.Field(i)
		queryTag := field.Tag.Get("query")

		// Если тег "query" существует и ключ есть в queryParams
		if queryTag != "" {
			if value, exists := queryParams[queryTag]; exists {
				destField := destValue.Field(i)

				// Если поле является вложенной структурой и содержит JSON
				if field.Type.Kind() == reflect.Struct && queryTag == "filters" {
					if err := json.Unmarshal([]byte(value), destField.Addr().Interface()); err != nil {
						log.Printf("Failed to parse JSON for field %s: %v\n", field.Name, err)
					}
					continue
				}

				// Преобразование значения в тип поля структуры
				if destField.CanSet() {
					switch destField.Kind() {
					case reflect.String:
						destField.SetString(value)
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						intValue, err := strconv.ParseInt(value, 10, 64)
						if err != nil {
							log.Printf("Failed to parse '%s' as int for field %s: %v\n", value, field.Name, err)
							continue
						}
						destField.SetInt(intValue)
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						uintValue, err := strconv.ParseUint(value, 10, 64)
						if err != nil {
							log.Printf("Failed to parse '%s' as uint for field %s: %v\n", value, field.Name, err)
							continue
						}
						destField.SetUint(uintValue)
					case reflect.Float32, reflect.Float64:
						floatValue, err := strconv.ParseFloat(value, 64)
						if err != nil {
							log.Printf("Failed to parse '%s' as float for field %s: %v\n", value, field.Name, err)
							continue
						}
						destField.SetFloat(floatValue)
					case reflect.Bool:
						boolValue, err := strconv.ParseBool(value)
						if err != nil {
							log.Printf("Failed to parse '%s' as bool for field %s: %v\n", value, field.Name, err)
							continue
						}
						destField.SetBool(boolValue)
					default:
						log.Printf("Unsupported type for field %s\n", field.Name)
					}
				}
			}
		}
	}
}
