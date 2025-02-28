package service

import (
	"context"
	"mime/multipart"

	"github.com/Fi44er/sdmedik/backend/module/category/domain"
	"github.com/Fi44er/sdmedik/backend/module/category/repository"
	file_service "github.com/Fi44er/sdmedik/backend/module/file/service"
	transaction_manager_repo "github.com/Fi44er/sdmedik/backend/module/transaction_manager/repository"
	"github.com/Fi44er/sdmedik/backend/shared/logger"
)

var _ ICategoryService = (*CategoryService)(nil)

type ICategoryService interface {
	Create(ctx context.Context, categoryDomain *domain.Category, files []*multipart.FileHeader) error
}

type CategoryService struct {
	logger *logger.Logger

	repo                   repository.ICategoryRepository
	transactionManagerRepo transaction_manager_repo.ITransactionManagerRepository
	fileServ               file_service.IFileService
}

func NewCategoryService(
	logger *logger.Logger,
	repo repository.ICategoryRepository,
	transactionManagerRepo transaction_manager_repo.ITransactionManagerRepository,
	fileServ file_service.IFileService,
) *CategoryService {
	return &CategoryService{
		logger:                 logger,
		repo:                   repo,
		transactionManagerRepo: transactionManagerRepo,
		fileServ:               fileServ,
	}
}
