package service

import (
	"context"
	"mime/multipart"

	file_service "github.com/Fi44er/sdmedik/backend/module/file/service"
	"github.com/Fi44er/sdmedik/backend/module/product/domain"
	"github.com/Fi44er/sdmedik/backend/module/product/repository"
	transaction_manager_repo "github.com/Fi44er/sdmedik/backend/module/transaction_manager/repository"
	events "github.com/Fi44er/sdmedik/backend/shared/eventbus"
	"github.com/Fi44er/sdmedik/backend/shared/logger"
)

var _ IProductService = (*ProductService)(nil)

type IProductService interface {
	Create(ctx context.Context, productDomain *domain.Product, files []*multipart.FileHeader) error

	GetByID(ctx context.Context, id string) (*domain.Product, error)
}

type ProductService struct {
	logger  *logger.Logger
	evenBus *events.EventBus

	repo                   repository.IProductRepository
	transactionManagerRepo transaction_manager_repo.ITransactionManagerRepository
	fileServ               file_service.IFileService
}

func NewProductService(
	logger *logger.Logger,
	evenBus *events.EventBus,
	repo repository.IProductRepository,
	transactionManagerRepo transaction_manager_repo.ITransactionManagerRepository,
	fileServ file_service.IFileService,
) *ProductService {
	return &ProductService{
		logger:                 logger,
		evenBus:                evenBus,
		repo:                   repo,
		transactionManagerRepo: transactionManagerRepo,
		fileServ:               fileServ,
	}
}
