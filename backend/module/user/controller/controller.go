package controller

import (
	"github.com/Fi44er/sdmedik/backend/module/user/service"
	"github.com/Fi44er/sdmedik/backend/shared/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var _ IUserController = (*UserController)(nil)

type IUserController interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error

	GetByID(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetMy(ctx *fiber.Ctx) error
}

type UserController struct {
	service service.IUserService

	logger    *logger.Logger
	validator *validator.Validate
}

func NewUserController(
	service service.IUserService,
	logger *logger.Logger,
	validator *validator.Validate,
) *UserController {
	return &UserController{
		service:   service,
		logger:    logger,
		validator: validator,
	}
}
