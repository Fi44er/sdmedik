package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/api/order"
	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	orderRepository "github.com/Fi44er/sdmedik/backend/internal/repository/order"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	orderService "github.com/Fi44er/sdmedik/backend/internal/service/order"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type OrderProvider struct {
	orderService    service.IOrderService
	orderImpl       *order.Implementation
	orderRepository repository.IOrderRepository

	logger    *logger.Logger
	validator *validator.Validate
	config    *config.Config
	db        *gorm.DB

	basketService  service.IBasketService
	certService    service.ICertificateService
	productService service.IProductService
}

func NewOrderProvider(
	logger *logger.Logger,
	validator *validator.Validate,
	db *gorm.DB,
	config *config.Config,
	basketService service.IBasketService,
	certService service.ICertificateService,
	productService service.IProductService,
) *OrderProvider {
	return &OrderProvider{
		logger:         logger,
		validator:      validator,
		db:             db,
		config:         config,
		basketService:  basketService,
		certService:    certService,
		productService: productService,
	}
}

func (p *OrderProvider) OrderRepository() repository.IOrderRepository {
	if p.orderRepository == nil {
		p.orderRepository = orderRepository.NewRepository(p.logger, p.db)
	}
	return p.orderRepository
}

func (p *OrderProvider) OrderService() service.IOrderService {
	if p.orderService == nil {
		p.orderService = orderService.NewService(p.logger, p.validator, p.config, p.OrderRepository(), p.basketService, p.certService, p.productService)
	}
	return p.orderService
}

func (p *OrderProvider) OrderImpl() *order.Implementation {
	if p.orderImpl == nil {
		p.orderImpl = order.NewImplementation(p.OrderService())
	}
	return p.orderImpl
}
