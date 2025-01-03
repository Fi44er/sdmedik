package constants

import "github.com/Fi44er/sdmedik/backend/pkg/errors"

var ErrCategoryNotFound = errors.New(404, "Category not found")
var ErrCharacteristicNotFound = errors.New(404, "Characteristic not found")
var ErrProductNotFound = errors.New(404, "Product not found")
var ErrUserNotFound = errors.New(404, "User not found")
