package repository

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/module/user/converter"
	"github.com/Fi44er/sdmedik/backend/module/user/domain"
	"github.com/Fi44er/sdmedik/backend/module/user/model"
	"gorm.io/gorm"
)

func (r *UserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	r.logger.Infof("Getting user by ID: %s...", id)
	user := new(model.User)
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Infof("User with ID %s not found", id)
			return nil, nil
		}

		r.logger.Errorf("Error getting user by ID: %v", err)
		return nil, err
	}

	userDomain := converter.ToDomainFromModel(user)
	return userDomain, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	r.logger.Infof("Getting user by email: %s...", email)
	userModel := new(model.User)
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(userModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Infof("User with email %s not found", email)
			return nil, nil
		}

		r.logger.Errorf("Error getting user by email: %v", err)
		return nil, err
	}

	userDomain := converter.ToDomainFromModel(userModel)
	return userDomain, nil
}

func (r *UserRepository) GetAll(ctx context.Context, offset, limit int) ([]domain.User, error) {
	r.logger.Infof("Getting all users...")
	usersModels := new([]model.User)
	if offset == 0 {
		offset = -1
	}

	if limit == 0 {
		limit = -1
	}
	if err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&usersModels).Error; err != nil {
		r.logger.Errorf("Error getting all users: %v", err)
		return nil, err
	}

	usersDomain := converter.ToDomainSliceFromModel(*usersModels)
	return usersDomain, nil
}
