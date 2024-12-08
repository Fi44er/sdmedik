package utils

import (
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/go-playground/validator/v10"
)

func CustomTypeValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String() // Получаем строковое значение
	switch model.Type(value) {
	case model.TypeString, model.TypeInt, model.TypeFloat, model.TypeBool:
		return true
	default:
		return false
	}
}
