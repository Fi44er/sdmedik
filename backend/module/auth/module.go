package module

import (
	"github.com/Fi44er/sdmedik/backend/config"
	"github.com/Fi44er/sdmedik/backend/module/auth/controller"
	"github.com/Fi44er/sdmedik/backend/module/auth/service"
	user_service "github.com/Fi44er/sdmedik/backend/module/user/service"
	"github.com/Fi44er/sdmedik/backend/shared/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type AuthModule struct {
	authServ       service.IAuthService
	authController controller.IAuthController
	logger         *logger.Logger
	validator      *validator.Validate
	cache          *redis.Client
	config         *config.Config

	userServ user_service.IUserService
}

func NewAuthModule(
	logger *logger.Logger,
	validator *validator.Validate,
	cache *redis.Client,
	config *config.Config,
	userServ user_service.IUserService,
) *AuthModule {
	return &AuthModule{
		logger:    logger,
		validator: validator,
		cache:     cache,
		config:    config,
		userServ:  userServ,
	}
}

func (m *AuthModule) AuthService() service.IAuthService {
	if m.authServ == nil {
		m.authServ = service.NewAuthService(m.logger, m.cache, m.config, m.userServ)
	}
	return m.authServ
}

func (m *AuthModule) AuthController() controller.IAuthController {
	if m.authController == nil {
		m.authController = controller.NewAuthController(m.logger, m.validator, m.AuthService(), m.config)
	}
	return m.authController
}

func (m *AuthModule) RegisterRoutes(router fiber.Router, middleware fiber.Handler) {
	auth := router.Group("/auth")
	auth.Post("/register", m.AuthController().Register)
	auth.Post("/login", m.AuthController().Login)
	auth.Post("/send-code", m.AuthController().SendCode)
	auth.Post("/verify-code", m.AuthController().VerifyCode)
	auth.Post("/refresh", middleware, m.AuthController().RefreshAccessToken)
	auth.Post("/logout", middleware, m.AuthController().Logout)
}
