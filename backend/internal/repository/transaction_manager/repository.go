package transactionmanager

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

type ITransactionManager interface {
	BeginTransaction(ctx context.Context) (*gorm.DB, error)
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB)
	WithTransaction(tx *gorm.DB) *gorm.DB
}

type Manager struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewManager(db *gorm.DB, logger *logger.Logger) *Manager {
	return &Manager{
		db:     db,
		logger: logger,
	}
}

func (m *Manager) BeginTransaction(ctx context.Context) (*gorm.DB, error) {
	m.logger.Info("Starting transaction...")
	tx := m.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		m.logger.Errorf("Failed to start transaction: %v", tx.Error)
		return nil, tx.Error
	}
	return tx, nil
}

func (m *Manager) Commit(tx *gorm.DB) error {
	m.logger.Info("Committing transaction...")
	if err := tx.Commit().Error; err != nil {
		m.logger.Errorf("Failed to commit transaction: %v", err)
		return err
	}
	return nil
}

func (m *Manager) Rollback(tx *gorm.DB) {
	m.logger.Warn("Rolling back transaction...")
	_ = tx.Rollback().Error
}

func (m *Manager) WithTransaction(tx *gorm.DB) *gorm.DB {
	return tx
}
