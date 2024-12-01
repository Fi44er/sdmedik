package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/api/auth"
	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	authService "github.com/Fi44er/sdmedik/backend/internal/service/auth"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
)

type AuthProvider struct {
	authService service.IAuthService
	authImpl    *auth.Implementation

	logger    *logger.Logger
	validator *validator.Validate
	config    *config.Config
	cache     *redis.Client

	userService service.IUserService
}

func NewAuthProvider(
	logger *logger.Logger,
	validator *validator.Validate,
	config *config.Config,
	cache *redis.Client,
	userService service.IUserService,
) *AuthProvider {
	return &AuthProvider{
		logger:      logger,
		validator:   validator,
		config:      config,
		cache:       cache,
		userService: userService,
	}
}

func (p *AuthProvider) AuthService() service.IAuthService {
	if p.authService == nil {
		p.authService = authService.NewService(p.logger, p.validator, p.config, p.cache, p.userService)
	}
	return p.authService
}

func (p *AuthProvider) AuthImpl() *auth.Implementation {
	if p.authImpl == nil {
		p.authImpl = auth.NewImplementation(p.AuthService(), p.config)
	}
	return p.authImpl
}
