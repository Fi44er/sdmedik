package auth

import (
	"context"
	"time"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) Login(ctx context.Context, dto *dto.Login, userAgent string) (string, string, error) {
	if err := s.validator.Struct(dto); err != nil {
		return "", "", errors.New(400, err.Error())
	}

	existUser, err := s.userService.GetByEmail(ctx, dto.Email)
	if err != nil {
		return "", "", err
	}

	if !utils.ComparePassword(existUser.Password, dto.Password) {
		s.logger.Info("Invalid email or password")
		return "", "", errors.New(400, "Invalid email or password")
	}

	accessTokenDetails, err := utils.CreateToken(existUser.ID, s.config.AccessTokenExpiresIn, s.config.AccessTokenPrivateKey)
	if err != nil {
		return "", "", errors.New(422, err.Error())
	}

	refreshTokenDetails, err := utils.CreateToken(existUser.ID, s.config.RefreshTokenExpiresIn, s.config.RefreshTokenPrivateKey)
	if err != nil {
		return "", "", errors.New(422, err.Error())
	}

	accessRedisKey := userAgent + ":" + accessTokenDetails.TokenUUID
	errAccess := s.cache.Set(ctx, accessRedisKey, existUser.ID, time.Unix(*accessTokenDetails.ExpiresIn, 0).Sub(time.Now())).Err()
	if errAccess != nil {
		return "", "", errors.New(422, errAccess.Error())
	}

	refreshRedisKey := userAgent + ":" + refreshTokenDetails.TokenUUID
	errRefresh := s.cache.Set(ctx, refreshRedisKey, existUser.ID, time.Unix(*refreshTokenDetails.ExpiresIn, 0).Sub(time.Now())).Err()
	if errRefresh != nil {
		return "", "", errors.New(422, errRefresh.Error())
	}

	return *accessTokenDetails.Token, *refreshTokenDetails.Token, nil
}
