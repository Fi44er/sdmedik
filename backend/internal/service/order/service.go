package order

import (
	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
)

var _ def.IOrderService = (*service)(nil)

type service struct {
	logger    *logger.Logger
	validator *validator.Validate
	config    *config.Config
	repo      repository.IOrderRepository

	basketService  def.IBasketService
	certService    def.ICertificateService
	productService def.IProductService
}

func NewService(
	logger *logger.Logger,
	validator *validator.Validate,
	config *config.Config,
	repo repository.IOrderRepository,
	basketService def.IBasketService,
	certService def.ICertificateService,
	productService def.IProductService,
) *service {
	return &service{
		logger:         logger,
		validator:      validator,
		config:         config,
		repo:           repo,
		basketService:  basketService,
		certService:    certService,
		productService: productService,
	}
}
