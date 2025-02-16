package product

import (
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	events "github.com/Fi44er/sdmedik/backend/pkg/evenbus"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
)

var _ def.IProductService = (*service)(nil)

type service struct {
	logger    *logger.Logger
	repo      repository.IProductRepository
	validator *validator.Validate
	evenBus   *events.EventBus

	categoryService            def.ICategoryService
	characteristicValueService def.ICharacteristicValueService
	transactionManagerRepo     repository.ITransactionManager
	imageService               def.IImageService
	characteristicService      def.ICharacteristicService
	certificateService         def.ICertificateService
}

func NewService(
	repo repository.IProductRepository,
	logger *logger.Logger,
	validator *validator.Validate,
	categoryService def.ICategoryService,
	characteristicValueService def.ICharacteristicValueService,
	transactionManagerRepo repository.ITransactionManager,
	imageService def.IImageService,
	characteristicService def.ICharacteristicService,
	evenBus *events.EventBus,
	certificateService def.ICertificateService,
) *service {
	return &service{
		repo:                       repo,
		logger:                     logger,
		validator:                  validator,
		categoryService:            categoryService,
		characteristicValueService: characteristicValueService,
		transactionManagerRepo:     transactionManagerRepo,
		imageService:               imageService,
		characteristicService:      characteristicService,
		evenBus:                    evenBus,
		certificateService:         certificateService,
	}
}
