package constant

import "github.com/Fi44er/sdmedik/backend/pkg/customerr"

var (
	ErrUserNotFound    = customerr.NewError(404, "user not found")
	ErrInvalidUserData = customerr.NewError(422, "invalid user data")

	ErrInternalServerError = customerr.NewError(500, "internal server error")
)
