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

func (s *AuthProvider) AuthService() service.IAuthService {
	if s.authService == nil {
		s.authService = authService.NewService(s.logger, s.validator, s.config, s.cache, s.userService)
	}
	return s.authService
}

func (s *AuthProvider) AuthImpl() *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.AuthService(), s.config)
	}
	return s.authImpl
}
