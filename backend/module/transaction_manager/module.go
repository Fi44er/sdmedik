package module

import (
	"github.com/Fi44er/sdmedik/backend/module/transaction_manager/repository"
	"github.com/Fi44er/sdmedik/backend/shared/logger"
	"gorm.io/gorm"
)

type TransactionManagerModule struct {
	repository repository.ITransactionManagerRepository
	logger     *logger.Logger
	db         *gorm.DB
}

func NewTransactionManagerModule(logger *logger.Logger, db *gorm.DB) *TransactionManagerModule {
	return &TransactionManagerModule{
		repository: repository.NewTransactionManagerRepository(db, logger),
		logger:     logger,
		db:         db,
	}
}

func (m *TransactionManagerModule) TransactionManagerRepository() repository.ITransactionManagerRepository {
	if m.repository == nil {
		m.repository = repository.NewTransactionManagerRepository(m.db, m.logger)
	}
	return m.repository
}
