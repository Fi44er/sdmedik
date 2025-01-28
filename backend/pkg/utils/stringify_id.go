package utils

import (
	"fmt"
	"strconv"
)

func StringifyID(id any) (string, error) {
	switch v := id.(type) {
	case string:
		return v, nil
	case int:
		return strconv.Itoa(v), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	default:
		return "", fmt.Errorf("тип поля 'ID' не поддерживается: %T", id)
	}
}
