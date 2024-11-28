package user

import (
	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
)

var _ def.IUserService = (*service)(nil)

type service struct {
	logger *logger.Logger
	repo   repository.IUserRepository

	validator *validator.Validate
	cache     *redis.Client
	config    *config.Config
}

func NewService(logger *logger.Logger, repo repository.IUserRepository, validator *validator.Validate, config *config.Config, cache *redis.Client) *service {
	return &service{
		logger:    logger,
		repo:      repo,
		validator: validator,
		config:    config,
		cache:     cache,
	}
}
