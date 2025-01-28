package characteristicvalue

import (
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
)

var _ def.ICharacteristicValueService = (*service)(nil)

type service struct {
	logger    *logger.Logger
	validator *validator.Validate
	repo      repository.ICharacteristicValueRepository
}

func NewService(
	logger *logger.Logger,
	validator *validator.Validate,
	repo repository.ICharacteristicValueRepository,
) *service {
	return &service{
		logger:    logger,
		validator: validator,
		repo:      repo,
	}
}
