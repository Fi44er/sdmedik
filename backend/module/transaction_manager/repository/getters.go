package repository

import (
	"context"

	"gorm.io/gorm"
)

func (m *TransactionManagerRepository) GetTransaction(ctx context.Context) (*gorm.DB, error) {
	tx, ok := ctx.Value(TransactionKey).(*gorm.DB)
	if !ok {
		return nil, gorm.ErrInvalidTransaction
	}
	return tx, nil
}
