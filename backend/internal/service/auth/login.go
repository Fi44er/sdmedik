package auth

import (
	"context"
	"time"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func (s *service) Login(ctx context.Context, data *dto.Login, userAgent string, session *session.Session) (string, string, error) {
	if err := s.validator.Struct(data); err != nil {
		return "", "", errors.New(400, err.Error())
	}

	existUser, err := s.userService.GetByEmail(ctx, data.Email)
	if err != nil {
		return "", "", err
	}

	if !utils.ComparePassword(existUser.Password, data.Password) {
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

	s.logger.Infof("session login: %+v", session)

	if err := s.basketService.Move(ctx, &dto.MoveBasket{
		UserID:  existUser.ID,
		Session: session,
	}); err != nil {
		return "", "", err
	}

	if session != nil {
		if err := session.Destroy(); err != nil {
			s.logger.Errorf("Error during destroying session: %s", err.Error())
		}
	}

	return *accessTokenDetails.Token, *refreshTokenDetails.Token, nil
}
