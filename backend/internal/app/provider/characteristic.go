package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	characteristicRepository "github.com/Fi44er/sdmedik/backend/internal/repository/characteristic"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	characteristicService "github.com/Fi44er/sdmedik/backend/internal/service/characteristic"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CharacteristicProvider struct {
	characteristicRepository repository.ICharacteristicRepository
	characteristicService    service.ICharacteristicService
	// characteristicImpl       *characteristic.Implementation

	logger    *logger.Logger
	db        *gorm.DB
	validator *validator.Validate
}

func NewCharacteristicProvider(
	logger *logger.Logger,
	db *gorm.DB,
	validator *validator.Validate,
) *CharacteristicProvider {
	return &CharacteristicProvider{
		logger:    logger,
		db:        db,
		validator: validator,
	}
}

func (p *CharacteristicProvider) CharacteristicRepository() repository.ICharacteristicRepository {
	if p.characteristicRepository == nil {
		p.characteristicRepository = characteristicRepository.NewRepository(p.logger, p.db)
	}

	return p.characteristicRepository
}

func (p *CharacteristicProvider) CharacteristicService() service.ICharacteristicService {
	if p.characteristicService == nil {
		p.characteristicService = characteristicService.NewService(p.logger, p.validator, p.CharacteristicRepository())
	}

	return p.characteristicService
}
