package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	transactionManagerRepository "github.com/Fi44er/sdmedik/backend/internal/repository/transaction_manager"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

type TransactionManagerProvider struct {
	transactionManagerRepository repository.ITransactionManager
	logger                       *logger.Logger
	db                           *gorm.DB
}

func NewTransactionManagerProvider(logger *logger.Logger, db *gorm.DB) *TransactionManagerProvider {
	return &TransactionManagerProvider{
		logger: logger,
		db:     db,
	}
}

func (p *TransactionManagerProvider) TransactionManager() repository.ITransactionManager {
	if p.transactionManagerRepository == nil {
		p.transactionManagerRepository = transactionManagerRepository.NewManager(p.db, p.logger)
	}

	return p.transactionManagerRepository
}
