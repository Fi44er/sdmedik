package auth

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) Logout(ctx context.Context, refreshToken string, accessTokenUUID string) error {
	if refreshToken == "" {
		return errors.New(403, "Token is invalid or session has expired")
	}

	tokenClaims, err := utils.ValidateToken(refreshToken, s.config.RefreshTokenPublicKey)
	if err != nil {
		return errors.New(403, err.Error())
	}

	_, err = s.cache.Del(ctx, tokenClaims.TokenUUID, accessTokenUUID).Result()
	if err != nil {
		return errors.New(502, err.Error())
	}

	return nil
}
