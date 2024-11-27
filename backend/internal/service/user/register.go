package user

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) Register(ctx context.Context, dto *dto.Register) error {
	s.logger.Info("Register user...")
	if err := s.validator.Struct(dto); err != nil {
		return errors.New(400, err.Error())
	}

	var user model.User
	existUser, err := s.repo.GetByEmail(ctx, dto.Email)
	if err != nil {
		return err
	}
	if existUser.ID != "" {
		return errors.New(409, "user with this email already exists")
	}

	dto.Password = utils.GeneratePassword(dto.Password)
	if err := utils.DtoToModel(dto, &user); err != nil {
		s.logger.Errorf("Error during conversion: %s", err.Error())
		return err
	}

	if err := s.repo.Create(ctx, &user); err != nil {
		return err
	}

	return nil
}
