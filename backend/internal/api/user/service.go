package user

import "github.com/Fi44er/sdmedik/backend/internal/service"

type Implementation struct {
	userService service.IUserService
}

func NewImplementation(userService service.IUserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
