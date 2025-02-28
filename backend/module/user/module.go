package module

import (
	controller "github.com/Fi44er/sdmedik/backend/module/user/controller"
	repository "github.com/Fi44er/sdmedik/backend/module/user/repository"
	service "github.com/Fi44er/sdmedik/backend/module/user/service"
	"github.com/Fi44er/sdmedik/backend/shared/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserModule struct {
	repository repository.IUserRepository
	controller controller.IUserController
	service    service.IUserService

	logger    *logger.Logger
	validator *validator.Validate
	db        *gorm.DB
}

func NewUserModule(logger *logger.Logger, validator *validator.Validate, db *gorm.DB) *UserModule {
	return &UserModule{
		logger:    logger,
		validator: validator,
		db:        db,
	}
}

func (m *UserModule) UserRepository() repository.IUserRepository {
	if m.repository == nil {
		m.repository = repository.NewUserRepository(m.logger, m.db)
	}
	return m.repository
}

func (m *UserModule) UserService() service.IUserService {
	if m.service == nil {
		m.service = service.NewUserService(m.UserRepository(), m.logger)
	}
	return m.service
}

func (m *UserModule) UserController() controller.IUserController {
	if m.controller == nil {
		m.controller = controller.NewUserController(m.UserService(), m.logger, m.validator)
	}
	return m.controller
}

func (m *UserModule) RegisterRoutes(router fiber.Router) {
	user := router.Group("/users")

	user.Post("/", m.UserController().Create)
	user.Put("/:id", m.UserController().Update)
	user.Delete("/:id", m.UserController().Delete)
	user.Get("/", m.UserController().GetAll)
	user.Get("/me", m.UserController().GetMy)
	user.Get("/:id", m.UserController().GetByID)
}
