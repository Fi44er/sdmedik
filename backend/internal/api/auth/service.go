package auth

import (
	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/internal/service"
)

type Implementation struct {
	authService service.IAuthService
	config      *config.Config
}

func NewImplementation(authService service.IAuthService, config *config.Config) *Implementation {
	return &Implementation{
		authService: authService,
		config:      config,
	}
}
