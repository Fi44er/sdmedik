package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/api/basket"
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	basketRepository "github.com/Fi44er/sdmedik/backend/internal/repository/basket"
	basketItemRepository "github.com/Fi44er/sdmedik/backend/internal/repository/basket_item"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	basketService "github.com/Fi44er/sdmedik/backend/internal/service/basket"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type BasketProvider struct {
	basketRepository     repository.IBasketRepository
	basketItemRepository repository.IBasketItemRepository
	basketService        service.IBasketService
	basketImpl           *basket.Implementation

	logger    *logger.Logger
	db        *gorm.DB
	validator *validator.Validate

	productService service.IProductService
}

func NewBasketProvider(
	logger *logger.Logger,
	db *gorm.DB,
	validator *validator.Validate,
	productService service.IProductService,
) *BasketProvider {
	return &BasketProvider{
		logger:         logger,
		db:             db,
		validator:      validator,
		productService: productService,
	}
}

func (p *BasketProvider) BasketRepository() repository.IBasketRepository {
	if p.basketRepository == nil {
		p.basketRepository = basketRepository.NewRepository(p.logger, p.db)
	}
	return p.basketRepository
}

func (p *BasketProvider) BasketItemRepository() repository.IBasketItemRepository {
	if p.basketItemRepository == nil {
		p.basketItemRepository = basketItemRepository.NewRepository(p.logger, p.db)
	}
	return p.basketItemRepository
}

func (p *BasketProvider) BasketService() service.IBasketService {
	if p.basketService == nil {
		p.basketService = basketService.NewService(
			p.logger,
			p.validator,
			p.BasketRepository(),
			p.productService,
			p.BasketItemRepository(),
		)
	}
	return p.basketService
}

func (p *BasketProvider) BasketImpl() *basket.Implementation {
	if p.basketImpl == nil {
		p.basketImpl = basket.NewImplementation(p.BasketService())
	}
	return p.basketImpl
}
