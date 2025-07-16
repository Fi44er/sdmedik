package page

import (
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
)

var _ def.IPageService = (*service)(nil)

type service struct {
	repo      repository.IPageRepository
	logger    *logger.Logger
	validator *validator.Validate
}

func NewService(
	repo repository.IPageRepository,
	logger *logger.Logger,
	validator *validator.Validate,
) *service {
	return &service{
		repo:      repo,
		logger:    logger,
		validator: validator,
	}
}
