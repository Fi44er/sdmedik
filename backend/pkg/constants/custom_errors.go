package constants

import "github.com/Fi44er/sdmedik/backend/pkg/errors"

var ErrCategoryNotFound = errors.New(404, "Category not found")

var ErrCharacteristicNotFound = errors.New(404, "Characteristic not found")

var ErrProductNotFound = errors.New(404, "Product not found")
var ErrProductWithArticleConflict = errors.New(409, "Product with this article already exists")

var ErrUserNotFound = errors.New(404, "User not found")

var ErrBasketNotFound = errors.New(404, "Basket not found")
var ErrBasketItemNotFound = errors.New(404, "Basket item not found")

var ErrCertificateNotFound = errors.New(404, "Certificate not found")

var ErrOrderNotFound = errors.New(404, "Order not found")
