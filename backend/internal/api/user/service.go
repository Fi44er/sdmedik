package user

import "github.com/Fi44er/sdmedik/backend/internal/service"

type Implementation struct {
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
