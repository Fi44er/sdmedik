package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/api/auth"
	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	authService "github.com/Fi44er/sdmedik/backend/internal/service/auth"
	events "github.com/Fi44er/sdmedik/backend/pkg/evenbus"
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
	eventBus  *events.EventBus

	userService   service.IUserService
	basketService service.IBasketService
}

func NewAuthProvider(
	logger *logger.Logger,
	validator *validator.Validate,
	config *config.Config,
	cache *redis.Client,
	eventBus *events.EventBus,
	userService service.IUserService,
	basketService service.IBasketService,
) *AuthProvider {
	return &AuthProvider{
		logger:        logger,
		validator:     validator,
		config:        config,
		cache:         cache,
		eventBus:      eventBus,
		userService:   userService,
		basketService: basketService,
	}
}

func (p *AuthProvider) AuthService() service.IAuthService {
	if p.authService == nil {
		serviceAuth, err := authService.NewService(p.logger, p.validator, p.config, p.cache, p.eventBus, p.userService, p.basketService)
		if err != nil {
			p.logger.Errorf("Error during initializing auth service: %s", err.Error())
			return nil
		}
		p.authService = serviceAuth
	}
	return p.authService
}

func (p *AuthProvider) AuthImpl() *auth.Implementation {
	if p.authImpl == nil {
		p.authImpl = auth.NewImplementation(p.AuthService(), p.config)
	}
	return p.authImpl
}
