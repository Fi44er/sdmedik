package constant

import "github.com/Fi44er/sdmedik/backend/pkg/customerr"

var (
	ErrUserNotFound       = customerr.NewError(404, "user not found")
	ErrUserAlreadyExists  = customerr.NewError(409, "user already exists")
	ErrInvalidPhoneNumber = customerr.NewError(422, "invalid phone number")

	ErrInvalidToken           = customerr.NewError(401, "invalid token")
	ErrCouldNotRefreshToken   = customerr.NewError(401, "could not refresh token")
	ErrAnauthorized           = customerr.NewError(401, "unauthorized")
	ErrForbidden              = customerr.NewError(403, "forbidden")
	ErrUnprocessableEntity    = customerr.NewError(422, "unprocessable entity")
	ErrInvalidEmailOrPassword = customerr.NewError(422, "invalid email or password")

	ErrInternalServerError = customerr.NewError(500, "internal server error")
)
