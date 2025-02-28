package module

import (
	"github.com/Fi44er/sdmedik/backend/config"
	"github.com/Fi44er/sdmedik/backend/module/file/repository"
	"github.com/Fi44er/sdmedik/backend/module/file/service"
	transaction_manager_repo "github.com/Fi44er/sdmedik/backend/module/transaction_manager/repository"
	"github.com/Fi44er/sdmedik/backend/shared/logger"
	"gorm.io/gorm"
)

type FileModule struct {
	repository repository.IFileRepository
	service    service.IFileService

	logger *logger.Logger
	config *config.Config
	db     *gorm.DB

	transactionManagerRepo transaction_manager_repo.ITransactionManagerRepository
}

func NewFileModule(
	logger *logger.Logger,
	config *config.Config,
	db *gorm.DB,
	transactionManagerRepo transaction_manager_repo.ITransactionManagerRepository,
) *FileModule {
	return &FileModule{
		logger:                 logger,
		config:                 config,
		db:                     db,
		transactionManagerRepo: transactionManagerRepo,
	}
}

func (m *FileModule) FileRepository() repository.IFileRepository {
	if m.repository == nil {
		m.repository = repository.NewFileRepository(m.logger, m.db)
	}
	return m.repository
}

func (m *FileModule) FileService() service.IFileService {
	if m.service == nil {
		m.service = service.NewFileService(m.logger, m.config, m.FileRepository(), m.transactionManagerRepo)
	}
	return m.service
}
