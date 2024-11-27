package user

import (
	"context"
	"os"
	"time"

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

	accessTokenDuration, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_EXPIRED_IN"))
	if err != nil {
		return "", err
	}

	accessTokenDetails, err := utils.CreateToken(existUser.ID, accessTokenDuration, os.Getenv("JWT_PRIVATE_KEY"))
	if err != nil {
		return "", err
	}

	refreshTokenDetails, err := utils.CreateToken()
	if err != nil {
		return "", err
	}

	errAccess := s.cache.Set(ctx, accessTokenDetails.TokenUUID, existUser.ID, 0).Err()

	// token, err := utils.GenerateToken(existUser.ID)
	// if err != nil {
	// 	s.logger.Errorf("Error during token generation: %s", err.Error())
	// 	return "", err
	// }
	//
	return "", nil
}
