package customerr

var (
	// --------------- AUTH ERRORS -----------------
	ErrInvalidEmailOrPassword = NewError(400, "Invalid email or password")
	ErrInvalidCode            = NewError(400, "Invalid code")
	ErrInvalidPhoneNumber     = NewError(400, "Invalid phone number")
	ErrInvalidToken           = NewError(401, "Invalid token")
	ErrCouldNotRefreshToken   = NewError(403, "Could not refresh token")

	// --------------- USER ERRORS -----------------
	ErrUserAlreadyExists = NewError(409, "User already exists")
	ErrUserNotFound      = NewError(404, "User not found")

	// --------------- PRODUCT ERRORS -----------------
	ErrProductNotFound = NewError(404, "Product not found")
)
