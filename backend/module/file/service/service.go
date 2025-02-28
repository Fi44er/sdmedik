package service

import (
	"context"
	"mime/multipart"

	"github.com/Fi44er/sdmedik/backend/config"
	"github.com/Fi44er/sdmedik/backend/module/file/domain"
	"github.com/Fi44er/sdmedik/backend/module/file/repository"
	transaction_manager_repo "github.com/Fi44er/sdmedik/backend/module/transaction_manager/repository"
	"github.com/Fi44er/sdmedik/backend/shared/logger"
)

var _ IFileService = (*FileService)(nil)

type IFileService interface {
	CreateMany(ctx context.Context, fileDomain *domain.File, files []*multipart.FileHeader) ([]string, error)
}

type FileService struct {
	repository             repository.IFileRepository
	logger                 *logger.Logger
	config                 *config.Config
	transactionManagerRepo transaction_manager_repo.ITransactionManagerRepository
}

func NewFileService(
	logger *logger.Logger,
	config *config.Config,
	repository repository.IFileRepository,
	transactionManagerRepo transaction_manager_repo.ITransactionManagerRepository,
) *FileService {
	return &FileService{
		logger:                 logger,
		config:                 config,
		repository:             repository,
		transactionManagerRepo: transactionManagerRepo,
	}
}
