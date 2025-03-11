package basket

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	events "github.com/Fi44er/sdmedik/backend/pkg/evenbus"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
)

var _ def.IBasketService = (*service)(nil)

type service struct {
	logger    *logger.Logger
	validator *validator.Validate
	repo      repository.IBasketRepository
	eventBus  *events.EventBus

	productService   def.IProductService
	basketItemRepo   repository.IBasketItemRepository
	promotionService def.IPromotionService
}

func NewService(
	logger *logger.Logger,
	validator *validator.Validate,
	repo repository.IBasketRepository,
	eventBus *events.EventBus,
	productService def.IProductService,
	basketItemRepo repository.IBasketItemRepository,
	promotionService def.IPromotionService,
) *service {
	svc := &service{
		logger:           logger,
		validator:        validator,
		repo:             repo,
		eventBus:         eventBus,
		productService:   productService,
		basketItemRepo:   basketItemRepo,
		promotionService: promotionService,
	}

	eventBus.Subscribe(events.EventDataMoveBasket, svc.handleDataMoveBasket)
	return svc
}

func (s *service) handleDataMoveBasket(event events.Event) {
	data := event.Data
	ctx := context.Background()
	dto, ok := data.(dto.MoveBasket)
	if !ok {
		return
	}
	if err := s.Move(ctx, &dto); err != nil {
		s.logger.Errorf("Ошибка при перемещении корзины: %v", err)
	}
}
