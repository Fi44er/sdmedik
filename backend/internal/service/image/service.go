package image

import (
	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
)

var _ def.IImageService = (*service)(nil)

type service struct {
	logger    *logger.Logger
	validator *validator.Validate
	repo      repository.IImageRepository
	config    *config.Config
}

func NewService(
	logger *logger.Logger,
	validator *validator.Validate,
	repo repository.IImageRepository,
	config *config.Config,
) *service {
	return &service{
		logger:    logger,
		validator: validator,
		repo:      repo,
		config:    config,
	}
}
