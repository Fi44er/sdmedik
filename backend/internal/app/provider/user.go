package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/api/user"
	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	userRepository "github.com/Fi44er/sdmedik/backend/internal/repository/user"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	userService "github.com/Fi44er/sdmedik/backend/internal/service/user"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserProvider struct {
	userRepository repository.IUserRepository
	userService    service.IUserService
	userImpl       *user.Implementation

	logger    *logger.Logger
	db        *gorm.DB
	validator *validator.Validate
	config    *config.Config
	cache     *redis.Client
}

func NewUserProvider(logger *logger.Logger, validator *validator.Validate, db *gorm.DB, config *config.Config, cache *redis.Client) *UserProvider {
	return &UserProvider{
		logger:    logger,
		validator: validator,
		db:        db,
		config:    config,
		cache:     cache,
	}
}

func (s *UserProvider) UserRepository() repository.IUserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.logger, s.db)
	}
	return s.userRepository
}

func (s *UserProvider) UserService() service.IUserService {
	if s.userService == nil {
		s.userService = userService.NewService(s.logger, s.UserRepository(), s.validator, s.config, s.cache)
	}

	return s.userService
}

func (s *UserProvider) UserImpl() *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(), s.config)
	}

	return s.userImpl
}
