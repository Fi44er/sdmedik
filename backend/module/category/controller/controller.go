package controller

import (
	"github.com/Fi44er/sdmedik/backend/module/category/service"
	"github.com/Fi44er/sdmedik/backend/shared/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var _ ICategoryController = (*CategoryController)(nil)

type ICategoryController interface {
	Create(ctx *fiber.Ctx) error
}

type CategoryController struct {
	logger    *logger.Logger
	validator *validator.Validate
	service   service.ICategoryService
}

func NewCategoryController(
	service service.ICategoryService,
	logger *logger.Logger,
	validator *validator.Validate,
) *CategoryController {
	return &CategoryController{
		service:   service,
		logger:    logger,
		validator: validator,
	}
}
