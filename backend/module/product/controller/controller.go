package controller

import (
	"github.com/Fi44er/sdmedik/backend/module/product/service"
	"github.com/Fi44er/sdmedik/backend/shared/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var _ IProductController = (*ProductController)(nil)

type IProductController interface {
	Create(ctx *fiber.Ctx) error

	GetByID(ctx *fiber.Ctx) error
}

type ProductController struct {
	service   service.IProductService
	logger    *logger.Logger
	validator *validator.Validate
}

func NewProductController(
	service service.IProductService,
	logger *logger.Logger,
	validator *validator.Validate,
) *ProductController {
	return &ProductController{
		service:   service,
		logger:    logger,
		validator: validator,
	}
}
