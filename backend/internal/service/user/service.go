package user

import (
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
)

var _ def.IUserService = (*service)(nil)

type service struct {
	logger    *logger.Logger
	repo      repository.IUserRepository
	validator validator.Validate
}

func NewService(logger *logger.Logger, repo repository.IUserRepository, validator validator.Validate) *service {
	return &service{
		logger:    logger,
		repo:      repo,
		validator: validator,
	}
}
