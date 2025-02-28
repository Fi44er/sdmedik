package controller

import (
	"github.com/Fi44er/sdmedik/backend/config"
	"github.com/Fi44er/sdmedik/backend/module/auth/service"
	"github.com/Fi44er/sdmedik/backend/shared/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var _ IAuthController = (*AuthController)(nil)

type IAuthController interface {
	SendCode(ctx *fiber.Ctx) error
	VerifyCode(ctx *fiber.Ctx) error
	RefreshAccessToken(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

type AuthController struct {
	logger    *logger.Logger
	validator *validator.Validate
	config    *config.Config

	authServ service.IAuthService
}

func NewAuthController(
	logger *logger.Logger,
	validator *validator.Validate,
	authServ service.IAuthService,
	config *config.Config,
) *AuthController {
	return &AuthController{
		logger:    logger,
		validator: validator,
		config:    config,
		authServ:  authServ,
	}
}
