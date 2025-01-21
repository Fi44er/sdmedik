package order

import (
	"github.com/Fi44er/sdmedik/backend/internal/config"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
)

var _ def.IOrderService = (*service)(nil)

type service struct {
	logger    *logger.Logger
	validator *validator.Validate
	config    *config.Config

	basketService def.IBasketService
	certService   def.ICertificateService
}

func NewService(
	logger *logger.Logger,
	validator *validator.Validate,
	config *config.Config,
	basketService def.IBasketService,
	certService def.ICertificateService,
) *service {
	return &service{
		logger:        logger,
		validator:     validator,
		config:        config,
		basketService: basketService,
		certService:   certService,
	}
}
