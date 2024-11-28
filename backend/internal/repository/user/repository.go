package user

import (
	"context"
	"fmt"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	def "github.com/Fi44er/sdmedik/backend/internal/repository"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

var _ def.IUserRepository = (*repository)(nil)

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

func (r *repository) Create(ctx context.Context, data *model.User) error {
	r.logger.Info("Creating user...")
	if err := r.db.WithContext(ctx).Create(data).Error; err != nil {
		r.logger.Errorf("Failed to create user: %v", err)
		return err
	}

	r.logger.Infof("User created successfully")
	return nil
}

func (r *repository) GetByID(ctx context.Context, id string) (model.User, error) {
	r.logger.Infof("Fetching user with ID: %s...", id)
	var user model.User
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Warnf("User with ID %s not found", id)
			return user, nil
		}
		r.logger.Errorf("Failed to fetch user with ID %s: %v", id, err)
		return model.User{}, err
	}
	r.logger.Info("User fetched successfully")
	return user, nil
}

func (r *repository) GetByEmail(ctx context.Context, email string) (model.User, error) {
	r.logger.Infof("Fetching user with email: %s...", email)
	var user model.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Warnf("User with email %s not found", email)
			return user, nil
		}
		r.logger.Errorf("Failed to fetch user with email %s: %v", email, err)
		return model.User{}, err
	}
	r.logger.Info("User fetched successfully")
	return user, nil
}

func (r *repository) GetAll(ctx context.Context, offset int, limit int) ([]model.User, error) {
	r.logger.Info("Fetching users...")
	var users []model.User
	if offset == 0 {
		offset = -1
	}

	if limit == 0 {
		limit = -1
	}

	if err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		r.logger.Errorf("Failed to fetch users: %v", err)
		return nil, err
	}
	r.logger.Info("Users fetched successfully")
	return users, nil
}

func (r *repository) Update(ctx context.Context, data *model.User) error {
	r.logger.Info("Updating user...")
	result := r.db.WithContext(ctx).Model(data).Updates(data)
	if err := result.Error; err != nil {
		r.logger.Errorf("Failed to update user: %v", err)
		return err
	}

	if result.RowsAffected == 0 {
		r.logger.Warnf("User with ID %s not found", data.ID)
		return fmt.Errorf("User not found")
	}

	r.logger.Info("User updated successfully")
	return nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	r.logger.Infof("Deleting user with ID: %s...", id)
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.User{})
	if err := result.Error; err != nil {
		r.logger.Errorf("Failed to delete user: %v", err)
		return err
	}

	if result.RowsAffected == 0 {
		r.logger.Warnf("User with ID %s not found", id)
		return fmt.Errorf("User not found")
	}

	r.logger.Infof("User deleted by ID: %v successfully", id)
	return nil
}
