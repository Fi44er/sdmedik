package order

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	def "github.com/Fi44er/sdmedik/backend/internal/repository"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

var _ def.IOrderRepository = (*repository)(nil)

type repository struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewRepository(
	logger *logger.Logger,
	db *gorm.DB,
) *repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) Create(ctx context.Context, data *model.Order) error {
	r.logger.Info("Creating order...")

	if err := r.db.WithContext(ctx).Create(data).Error; err != nil {
		r.logger.Errorf("Failed to create order: %v", err)
		return err
	}

	r.logger.Info("Order created successfully")
	return nil
}

func (r *repository) AddItems(ctx context.Context, items *[]model.OrderItem) error {
	r.logger.Info("Adding order items...")

	if err := r.db.WithContext(ctx).Create(items).Error; err != nil {
		r.logger.Errorf("Failed to add order items: %v", err)
		return err
	}

	r.logger.Info("Order items added successfully")
	return nil
}

func (r *repository) GetAll(ctx context.Context, offset int, limit int) (*[]model.Order, error) {
	r.logger.Info("Fetching orders...")
	orders := new([]model.Order)
	request := r.db.WithContext(ctx).Preload("Items")
	request = request.Order("created_at DESC")
	if offset != 0 {
		request = request.Offset(offset)
	}

	if limit != 0 {
		request = request.Limit(limit)
	}
	if err := request.Find(&orders).Error; err != nil {
		r.logger.Errorf("Failed to fetch orders: %v", err)
		return nil, err
	}
	r.logger.Info("Orders fetched successfully")
	return orders, nil
}
