package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/api/promotion"
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	promotionRepository "github.com/Fi44er/sdmedik/backend/internal/repository/promotion"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	promotionService "github.com/Fi44er/sdmedik/backend/internal/service/promotion"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type PromotionProvider struct {
	promotionService    service.IPromotionService
	promotionRepository repository.IPromotionRepository
	promotionImpl       *promotion.Implementation

	logger    *logger.Logger
	db        *gorm.DB
	validator *validator.Validate
}

func NewPromotionProvider(
	logger *logger.Logger,
	db *gorm.DB,
	validator *validator.Validate,
) *PromotionProvider {
	return &PromotionProvider{
		logger:    logger,
		db:        db,
		validator: validator,
	}
}

func (p *PromotionProvider) PromotionRepository() repository.IPromotionRepository {
	if p.promotionRepository == nil {
		p.promotionRepository = promotionRepository.NewRepository(p.logger, p.db)
	}
	return p.promotionRepository
}

func (p *PromotionProvider) PromotionService() service.IPromotionService {
	if p.promotionService == nil {
		p.promotionService = promotionService.NewService(p.PromotionRepository(), p.logger, p.validator)
	}
	return p.promotionService
}

func (p *PromotionProvider) PromotionImpl() *promotion.Implementation {
	if p.promotionImpl == nil {
		p.promotionImpl = promotion.NewImplementation(p.PromotionService())
	}
	return p.promotionImpl
}
