package auth

import (
	"github.com/Fi44er/sdmedik/backend/internal/config"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
)

var _ def.IAuthService = (*service)(nil)

type service struct {
	logger    *logger.Logger
	validator *validator.Validate
	cache     *redis.Client
	config    *config.Config

	userService def.IUserService
}

func NewService(
	logger *logger.Logger,
	validator *validator.Validate,
	config *config.Config,
	cache *redis.Client,
	userService def.IUserService,
) *service {
	return &service{
		logger:      logger,
		validator:   validator,
		config:      config,
		cache:       cache,
		userService: userService,
	}
}
