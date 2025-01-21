package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/api/order"
	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	orderService "github.com/Fi44er/sdmedik/backend/internal/service/order"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
)

type OrderProvider struct {
	orderService service.IOrderService
	orderImpl    *order.Implementation

	logger    *logger.Logger
	validator *validator.Validate
	config    *config.Config

	basketService service.IBasketService
	certService   service.ICertificateService
}

func NewOrderProvider(
	logger *logger.Logger,
	validator *validator.Validate,
	config *config.Config,
	basketService service.IBasketService,
	certService service.ICertificateService,
) *OrderProvider {
	return &OrderProvider{
		logger:        logger,
		validator:     validator,
		config:        config,
		basketService: basketService,
		certService:   certService,
	}
}

func (p *OrderProvider) OrderService() service.IOrderService {
	if p.orderService == nil {
		p.orderService = orderService.NewService(p.logger, p.validator, p.config, p.basketService, p.certService)
	}
	return p.orderService
}

func (p *OrderProvider) OrderImpl() *order.Implementation {
	if p.orderImpl == nil {
		p.orderImpl = order.NewImplementation(p.OrderService())
	}
	return p.orderImpl
}
