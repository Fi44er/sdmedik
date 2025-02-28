package repository

import (
	"context"

	"gorm.io/gorm"
)

func (r *TransactionManagerRepository) Begin(ctx context.Context) (*gorm.DB, context.Context, error) {
	r.logger.Info("Starting transaction...")
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		r.logger.Errorf("✖ Failed to start transaction: %v", tx.Error)
		return nil, nil, tx.Error
	}
	r.logger.Info("Transaction started successfully")
	ctx = context.WithValue(ctx, TransactionKey, tx)
	return tx, ctx, nil
}

func (r *TransactionManagerRepository) Commit(ctx context.Context, tx *gorm.DB) error {
	r.logger.Info("Committing transaction...")
	err := tx.Commit().Error
	if err != nil {
		r.logger.Errorf("✖ Failed to commit transaction: %v", err)
		return err
	}
	r.logger.Info("Transaction committed successfully")
	return nil
}

func (r *TransactionManagerRepository) Rollback(ctx context.Context, tx *gorm.DB) error {
	r.logger.Info("Rolling back transaction...")
	err := tx.Rollback().Error
	if err != nil {
		r.logger.Errorf("✖ Failed to rollback transaction: %v", err)
		return err
	}
	r.logger.Info("Transaction rolled back successfully")
	return nil
}
