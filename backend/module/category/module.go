package module

import (
	"github.com/Fi44er/sdmedik/backend/module/category/controller"
	"github.com/Fi44er/sdmedik/backend/module/category/repository"
	"github.com/Fi44er/sdmedik/backend/module/category/service"
	file_repository "github.com/Fi44er/sdmedik/backend/module/file/repository"
	file_service "github.com/Fi44er/sdmedik/backend/module/file/service"
	transaction_manager_repo "github.com/Fi44er/sdmedik/backend/module/transaction_manager/repository"
	"github.com/Fi44er/sdmedik/backend/shared/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CategoryModule struct {
	repository repository.ICategoryRepository
	service    service.ICategoryService
	controller controller.ICategoryController

	logger    *logger.Logger
	validator *validator.Validate
	db        *gorm.DB

	transactionManagerRepo transaction_manager_repo.ITransactionManagerRepository
	fileServ               file_service.IFileService
	fileRepo               file_repository.IFileRepository
}

func NewCategoryModule(
	logger *logger.Logger,
	validator *validator.Validate,
	db *gorm.DB,
	transactionManagerRepo transaction_manager_repo.ITransactionManagerRepository,
	fileServ file_service.IFileService,
	fileRepo file_repository.IFileRepository,
) *CategoryModule {
	return &CategoryModule{
		logger:                 logger,
		validator:              validator,
		db:                     db,
		transactionManagerRepo: transactionManagerRepo,
		fileServ:               fileServ,
		fileRepo:               fileRepo,
	}
}

func (m *CategoryModule) CategoryRepository() repository.ICategoryRepository {
	if m.repository == nil {
		m.repository = repository.NewCategoryRepository(m.logger, m.db, m.fileRepo)
	}
	return m.repository
}

func (m *CategoryModule) CategoryService() service.ICategoryService {
	if m.service == nil {
		m.service = service.NewCategoryService(m.logger, m.CategoryRepository(), m.transactionManagerRepo, m.fileServ)
	}
	return m.service
}

func (m *CategoryModule) CategoryController() controller.ICategoryController {
	if m.controller == nil {
		m.controller = controller.NewCategoryController(m.CategoryService(), m.logger, m.validator)
	}
	return m.controller
}

func (m *CategoryModule) RegisterRoutes(router fiber.Router) {
	category := router.Group("/categories")

	category.Post("/", m.CategoryController().Create)
	category.Get("/:id", m.CategoryController().GetByID)
	category.Get("/", m.CategoryController().GetAll)
}
