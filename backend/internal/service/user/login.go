package user

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) Login(ctx context.Context, dto *dto.Login) (string, error) {
	if err := s.validator.Struct(dto); err != nil {
		return "", errors.New(400, err.Error())
	}

	existUser, err := s.repo.GetByEmail(ctx, dto.Email)
	if err != nil {
		return "", err
	}
	if existUser.ID == "" {
		return "", errors.New(404, "User not found")
	}

	if !utils.ComparePassword(existUser.Password, dto.Password) {
		s.logger.Info("Invalid email or password")
		return "", errors.New(400, "Invalid email or password")
	}

	token, err := utils.GenerateToken(existUser.ID)
	if err != nil {
		s.logger.Errorf("Error during token generation: %s", err.Error())
		return "", err
	}

	return token, nil
}
