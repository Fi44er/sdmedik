package user

import (
	user_handler "github.com/Fi44er/sdmedik/backend/internal/module/user/delivery/http/user"
	user_repository "github.com/Fi44er/sdmedik/backend/internal/module/user/infrastructure/repository/user"
	user_usecase "github.com/Fi44er/sdmedik/backend/internal/module/user/usecase/user"
	"github.com/gofiber/fiber/v2"

	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserModule struct {
	repository *user_repository.UserRepository
	useCase    *user_usecase.UserUsecase
	handler    *user_handler.UserHandler

	logger    *logger.Logger
	validator *validator.Validate
	db        *gorm.DB
	redis     *redis.Client
}

func NewUserModule(
	logger *logger.Logger,
	validator *validator.Validate,
	db *gorm.DB,
	redis *redis.Client,
) *UserModule {
	return &UserModule{
		logger:    logger,
		validator: validator,
		db:        db,
		redis:     redis,
	}
}

func (m *UserModule) Init() {
	m.repository = user_repository.NewUserRepository(m.logger, m.db)
	m.useCase = user_usecase.NewUserUsecase(m.repository, m.logger)
	m.handler = user_handler.NewUserHandler(m.useCase, m.logger, m.validator)
}

func (m *UserModule) InitDelivery(router fiber.Router) {
	m.handler.RegisterRoutes(router)
}
