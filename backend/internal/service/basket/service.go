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

	userService    def.IUserService
	productService def.IProductService
	basketItemRepo repository.IBasketItemRepository
}

func NewService(
	logger *logger.Logger,
	validator *validator.Validate,
	repo repository.IBasketRepository,
	userService def.IUserService,
	productService def.IProductService,
	basketItemRepo repository.IBasketItemRepository,
) *service {
	return &service{
		logger:         logger,
		validator:      validator,
		repo:           repo,
		userService:    userService,
		productService: productService,
		basketItemRepo: basketItemRepo,
	}
}
