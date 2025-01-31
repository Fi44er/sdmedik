package auth

import (
	"context"
	"time"

	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
	"github.com/redis/go-redis/v9"
)

func (s *service) RefreshAccessToken(ctx context.Context, refreshToken string, userAgent string) (string, error) {
	if refreshToken == "" {
		return "", errors.New(403, "Could not refresh access token")
	}

	tokenClaims, err := utils.ValidateToken(refreshToken, s.config.RefreshTokenPublicKey)
	if err != nil {
		return "", errors.New(403, err.Error())
	}

	refreshRedisKey := userAgent + ":" + tokenClaims.TokenUUID
	userID, err := s.cache.Get(ctx, refreshRedisKey).Result()
	if err == redis.Nil {
		return "", errors.New(403, "Could not refresh access token")
	}

	user, err := s.userService.GetByID(ctx, userID)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New(404, "User not found")
	}

	accessTokenDetails, err := utils.CreateToken(user.ID, s.config.AccessTokenExpiresIn, s.config.AccessTokenPrivateKey)
	if err != nil {
		return "", errors.New(422, err.Error())
	}

	accessRedisKey := userAgent + ":" + accessTokenDetails.TokenUUID
	errAccess := s.cache.Set(ctx, accessRedisKey, user.ID, time.Unix(*accessTokenDetails.ExpiresIn, 0).Sub(time.Now())).Err()
	if errAccess != nil {
		return "", errors.New(422, errAccess.Error())
	}

	return *accessTokenDetails.Token, nil
}
