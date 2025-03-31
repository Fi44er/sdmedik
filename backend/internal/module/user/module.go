package user

import (
	"time"

	"github.com/Fi44er/sdmedik/backend/internal/config"
	user_handler "github.com/Fi44er/sdmedik/backend/internal/module/user/delivery/http/user"
	user_repository "github.com/Fi44er/sdmedik/backend/internal/module/user/infrastructure/repository/user"
	auth_usecase "github.com/Fi44er/sdmedik/backend/internal/module/user/usecase/auth"
	user_usecase "github.com/Fi44er/sdmedik/backend/internal/module/user/usecase/user"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"

	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	redis_manager "github.com/Fi44er/sdmedik/backend/pkg/redis"
	"github.com/Fi44er/sdmedik/backend/pkg/session"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserModule struct {
	userRepository *user_repository.UserRepository
	userUsecase    *user_usecase.UserUsecase
	userHandler    *user_handler.UserHandler

	authUsecase *auth_usecase.AuthUsecase

	logger         *logger.Logger
	validator      *validator.Validate
	db             *gorm.DB
	redisManager   redis_manager.IRedisManager
	redis          *redis.Client
	sessionManager *session.SessionManager
	config         *config.Config
}

func NewUserModule(
	logger *logger.Logger,
	validator *validator.Validate,
	db *gorm.DB,
	redisManager redis_manager.IRedisManager,
	redis *redis.Client,
) *UserModule {
	return &UserModule{
		logger:       logger,
		validator:    validator,
		db:           db,
		redisManager: redisManager,
		sessionManager: session.NewSessionManager(
			session.NewRedisSessionStore(redis),
			30*time.Minute,
			1*time.Hour,
			12*time.Hour,
			"session",
		),
	}
}

func (m *UserModule) Init() {
	m.userRepository = user_repository.NewUserRepository(m.logger, m.db)
	m.userUsecase = user_usecase.NewUserUsecase(m.userRepository, m.logger)
	m.userHandler = user_handler.NewUserHandler(m.userUsecase, m.logger, m.validator)

	m.authUsecase = auth_usecase.NewAuthUsecase(m.logger, m.redisManager, m.config, m.userUsecase)
}

func (m *UserModule) InitDelivery(router fiber.Router) {
	m.userHandler.RegisterRoutes(router)
}
