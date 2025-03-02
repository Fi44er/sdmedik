package module

import (
	file_repository "github.com/Fi44er/sdmedik/backend/module/file/repository"
	file_service "github.com/Fi44er/sdmedik/backend/module/file/service"
	"github.com/Fi44er/sdmedik/backend/module/product/controller"
	"github.com/Fi44er/sdmedik/backend/module/product/repository"
	"github.com/Fi44er/sdmedik/backend/module/product/service"
	transaction_manager_repo "github.com/Fi44er/sdmedik/backend/module/transaction_manager/repository"
	events "github.com/Fi44er/sdmedik/backend/shared/eventbus"
	"github.com/Fi44er/sdmedik/backend/shared/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductModule struct {
	repository repository.IProductRepository
	service    service.IProductService
	controller controller.IProductController

	logger    *logger.Logger
	validator *validator.Validate
	db        *gorm.DB
	evenBus   *events.EventBus

	transactionMaanagerRepo transaction_manager_repo.ITransactionManagerRepository
	fileServ                file_service.IFileService
	fileRepo                file_repository.IFileRepository
}

func NewProductModule(
	logger *logger.Logger,
	validator *validator.Validate,
	db *gorm.DB,
	evenBus *events.EventBus,
	transactionMaanagerRepo transaction_manager_repo.ITransactionManagerRepository,
	fileServ file_service.IFileService,
	fileRepo file_repository.IFileRepository,
) *ProductModule {
	return &ProductModule{
		logger:                  logger,
		validator:               validator,
		db:                      db,
		evenBus:                 evenBus,
		transactionMaanagerRepo: transactionMaanagerRepo,
		fileServ:                fileServ,
		fileRepo:                fileRepo,
	}
}

func (m *ProductModule) ProductRepository() repository.IProductRepository {
	if m.repository == nil {
		m.repository = repository.NewProductRepository(m.logger, m.db, m.fileRepo)
	}
	return m.repository
}

func (m *ProductModule) ProductService() service.IProductService {
	if m.service == nil {
		m.service = service.NewProductService(m.logger, m.evenBus, m.ProductRepository(), m.transactionMaanagerRepo, m.fileServ)
	}
	return m.service
}

func (m *ProductModule) ProductController() controller.IProductController {
	if m.controller == nil {
		m.controller = controller.NewProductController(m.ProductService(), m.logger, m.validator)
	}
	return m.controller
}

func (m *ProductModule) RegisterRoutes(router fiber.Router) {
	product := router.Group("/products")

	product.Post("/", m.ProductController().Create)
	product.Get("/:id", m.ProductController().GetByID)
	product.Get("/", m.ProductController().GetAll)
}
