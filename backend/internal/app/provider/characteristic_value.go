package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	characteristicValueRepository "github.com/Fi44er/sdmedik/backend/internal/repository/characteristic_value"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	characteristicValueService "github.com/Fi44er/sdmedik/backend/internal/service/characteristic_value"
)

type CharacteristicValueProvider struct {
	characteristicValueRepository repository.ICharacteristicValueRepository
	characteristicValueService    service.ICharacteristicValueService

	logger    *logger.Logger
	db        *gorm.DB
	validator *validator.Validate
}

func NewChracteristicValueProvider(
	logger *logger.Logger,
	db *gorm.DB,
	validator *validator.Validate,
) *CharacteristicValueProvider {
	return &CharacteristicValueProvider{
		logger:    logger,
		db:        db,
		validator: validator,
	}
}

func (p *CharacteristicValueProvider) CharacteristicValueRepository() repository.ICharacteristicValueRepository {
	if p.characteristicValueRepository == nil {
		p.characteristicValueRepository = characteristicValueRepository.NewRepository(p.logger, p.db)
	}
	return p.characteristicValueRepository
}

func (p *CharacteristicValueProvider) CharacteristicValueService() service.ICharacteristicValueService {
	if p.characteristicValueService == nil {
		p.characteristicValueService = characteristicValueService.NewService(p.logger, p.validator, p.CharacteristicValueRepository())
	}
	return p.characteristicValueService
}
