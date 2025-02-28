package repository

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/shared/logger"
	"gorm.io/gorm"
)

const TransactionKey = "transaction"

var _ ITransactionManagerRepository = (*TransactionManagerRepository)(nil)

type ITransactionManagerRepository interface {
	Begin(ctx context.Context) (*gorm.DB, context.Context, error)
	Commit(ctx context.Context, tx *gorm.DB) error
	Rollback(ctx context.Context, tx *gorm.DB) error
	GetTransaction(ctx context.Context) (*gorm.DB, error)
}

type TransactionManagerRepository struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewTransactionManagerRepository(db *gorm.DB, logger *logger.Logger) *TransactionManagerRepository {
	return &TransactionManagerRepository{
		db:     db,
		logger: logger,
	}
}
