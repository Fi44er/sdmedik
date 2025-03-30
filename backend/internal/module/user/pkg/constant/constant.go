package constant

import "github.com/Fi44er/sdmedik/backend/pkg/customerr"

var (
	ErrUserNotFound = customerr.NewError(404, "user not found")
)
