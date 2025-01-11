package basket

import (
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
)

var _ def.IBasketService = (*service)(nil)

type service struct {
	logger    *logger.Logger
	validator *validator.Validate

	userService    def.IUserService
	productService def.IProductService
}

func NewService(
	logger *logger.Logger,
	validator *validator.Validate,
	userService def.IUserService,
	productService def.IProductService,
) *service {
	return &service{
		logger:         logger,
		validator:      validator,
		userService:    userService,
		productService: productService,
	}
}
