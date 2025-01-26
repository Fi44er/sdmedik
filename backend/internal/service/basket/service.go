package basket

import (
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
)

var _ def.IBasketService = (*service)(nil)

type service struct {
	logger    *logger.Logger
	validator *validator.Validate
	repo      repository.IBasketRepository

	productService   def.IProductService
	basketItemRepo   repository.IBasketItemRepository
	promotionService def.IPromotionService
}

func NewService(
	logger *logger.Logger,
	validator *validator.Validate,
	repo repository.IBasketRepository,
	productService def.IProductService,
	basketItemRepo repository.IBasketItemRepository,
	promotionService def.IPromotionService,
) *service {
	return &service{
		logger:           logger,
		validator:        validator,
		repo:             repo,
		productService:   productService,
		basketItemRepo:   basketItemRepo,
		promotionService: promotionService,
	}
}
