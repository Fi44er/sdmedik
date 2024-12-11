package product

import (
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
)

var _ def.IProductService = (*service)(nil)

type service struct {
	logger    *logger.Logger
	repo      repository.IProductRepository
	validator *validator.Validate

	categoryService            def.ICategoryService
	characteristicValueService def.ICharacteristicValueService
	transactionManagerRepo     repository.ITransactionManager
}

func NewService(
	repo repository.IProductRepository,
	logger *logger.Logger,
	validator *validator.Validate,
	categoryService def.ICategoryService,
	characteristicValueService def.ICharacteristicValueService,
	transactionManagerRepo repository.ITransactionManager,
) *service {
	return &service{
		repo:                       repo,
		logger:                     logger,
		validator:                  validator,
		categoryService:            categoryService,
		characteristicValueService: characteristicValueService,
		transactionManagerRepo:     transactionManagerRepo,
	}
}
