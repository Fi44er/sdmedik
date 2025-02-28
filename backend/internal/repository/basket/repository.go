package basket

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	def "github.com/Fi44er/sdmedik/backend/internal/repository"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

var _ def.IBasketRepository = (*repository)(nil)

type repository struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewRepository(logger *logger.Logger, db *gorm.DB) *repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) Create(ctx context.Context, data *model.Basket) error {
	r.logger.Info("Creating basket...")
	if err := r.db.WithContext(ctx).Create(data).Error; err != nil {
		r.logger.Errorf("Failed to create basket: %v", err)
		return err
	}
	r.logger.Infof("Basket created successfully")
	return nil
}

func (r *repository) GetByUserID(ctx context.Context, userID string) (*model.Basket, error) {
	r.logger.Info("Fetching basket by userID...")
	basket := new(model.Basket)
	if err := r.db.WithContext(ctx).Preload("Items").Where("user_id = ?", userID).Find(basket, "user_id = ?", userID).Error; err != nil {
		r.logger.Errorf("Failed to fetch basket by userID: %v", err)
		return nil, err
	}
	r.logger.Info("Basket fetched by userID successfully")
	return basket, nil
}

func (r *repository) GetGuestBasketByID(ctx context.Context, id string) (*model.GuestBasket, error) {
	r.logger.Info("Fetching guest basket by ID...")
	basket := new(model.GuestBasket)
	if err := r.db.WithContext(ctx).Preload("Items").Where("id = ?", id).Find(basket, "id = ?", id).Error; err != nil {
		r.logger.Errorf("Failed to fetch guest basket by ID: %v", err)
		return nil, err
	}
	r.logger.Info("Guest basket fetched by ID successfully")
	return basket, nil
}

func (r *repository) CreateGuestBasket(ctx context.Context, basket *model.GuestBasket) error {
	r.logger.Info("Creating guest basket...")
	if err := r.db.WithContext(ctx).Create(basket).Error; err != nil {
		r.logger.Errorf("Failed to create guest basket: %v", err)
		return err
	}
	r.logger.Info("Guest basket created successfully")
	return nil
}
