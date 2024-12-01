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

func NewUserProvider(
	logger *logger.Logger,
	validator *validator.Validate,
	db *gorm.DB,
	config *config.Config,
	cache *redis.Client,
) *UserProvider {
	return &UserProvider{
		logger:    logger,
		validator: validator,
		db:        db,
		config:    config,
		cache:     cache,
	}
}

func (p *UserProvider) UserRepository() repository.IUserRepository {
	if p.userRepository == nil {
		p.userRepository = userRepository.NewRepository(p.logger, p.db)
	}
	return p.userRepository
}

func (p *UserProvider) UserService() service.IUserService {
	if p.userService == nil {
		p.userService = userService.NewService(p.logger, p.UserRepository(), p.validator, p.config, p.cache)
	}

	return p.userService
}

func (p *UserProvider) UserImpl() *user.Implementation {
	if p.userImpl == nil {
		p.userImpl = user.NewImplementation(p.UserService(), p.config)
	}

	return p.userImpl
}
