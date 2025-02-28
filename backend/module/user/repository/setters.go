package repository

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/module/user/converter"
	"github.com/Fi44er/sdmedik/backend/module/user/domain"
	"github.com/Fi44er/sdmedik/backend/module/user/model"
)

func (r *UserRepository) Create(ctx context.Context, userDomain *domain.User) error {
	r.logger.Infof("Creating user %+v...", userDomain)
	userModel := converter.ToModelFromDomain(userDomain)
	if err := r.db.WithContext(ctx).Create(userModel).Error; err != nil {
		r.logger.Errorf("Error creating user: %v", err)
		return err
	}
	r.logger.Info("User created successfully")
	return nil
}

func (r *UserRepository) Update(ctx context.Context, userDomain *domain.User) error {
	r.logger.Info("Updating user...")
	userModel := converter.ToModelFromDomain(userDomain)
	if err := r.db.WithContext(ctx).Model(userModel).Where("id = ?", userModel.ID).Updates(userModel).Error; err != nil {
		r.logger.Errorf("Error updating user: %v", err)
		return err
	}
	r.logger.Info("User updated successfully")
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	r.logger.Info("Deleting user...")
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		r.logger.Errorf("Error deleting user: %v", err)
		return err
	}
	r.logger.Info("User deleted successfully")
	return nil
}
