package user

import (
	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/internal/service"
)

type Implementation struct {
	userService service.IUserService
	config      *config.Config
}

func NewImplementation(userService service.IUserService, config *config.Config) *Implementation {
	return &Implementation{
		userService: userService,
		config:      config,
	}
}
