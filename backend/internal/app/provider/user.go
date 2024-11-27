package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/api/user"
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	userRepository "github.com/Fi44er/sdmedik/backend/internal/repository/user"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	userService "github.com/Fi44er/sdmedik/backend/internal/service/user"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserProvider struct {
	userRepository repository.IUserRepository
	userService    service.IUserService
	userImpl       *user.Implementation
	logger         *logger.Logger
	db             *gorm.DB
	validator      *validator.Validate
}

func NewUserProvider(logger *logger.Logger, vavalidator *validator.Validate) *UserProvider {
	return &UserProvider{
		logger:    logger,
		validator: vavalidator,
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
		s.userService = userService.NewService(s.logger, s.UserRepository(), s.validator)
	}

	return s.userService
}

func (s *UserProvider) UserImpl() *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService())
	}

	return s.userImpl
}
