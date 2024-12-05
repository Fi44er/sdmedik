package category

import (
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
)

var _ def.ICategoryService = (*service)(nil)

type service struct {
	logger    *logger.Logger
	repo      repository.ICategoryRepository
	validator *validator.Validate

	characteristicService def.ICharacteristicService
}

func NewService(
	repo repository.ICategoryRepository,
	logger *logger.Logger,
	validator *validator.Validate,
	characteristicService def.ICharacteristicService,
) *service {
	return &service{
		repo:                  repo,
		logger:                logger,
		validator:             validator,
		characteristicService: characteristicService,
	}
}
