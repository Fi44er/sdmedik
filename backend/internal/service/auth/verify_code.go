package auth

import (
	"context"
	"encoding/json"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) VerifyCode(ctx context.Context, data *dto.VerifyCode) error {
	hashEmail, err := utils.HashString(data.Email)
	if err != nil {
		s.logger.Errorf("Error during generating hash: %s", err.Error())
		return err
	}
	codeFromCache, err := s.cache.Get(ctx, "verification_codes_"+hashEmail).Result()
	if err != nil {
		s.logger.Errorf("Error during getting verification code: %s", err.Error())
		return errors.New(422, err.Error())
	}

	if codeFromCache != data.Code {
		s.logger.Info("Invalid verification code")
		return errors.New(400, "Invalid verification code")
	}

	if err := s.cache.Del(ctx, "verification_codes_"+hashEmail).Err(); err != nil {
		s.logger.Errorf("Error during deleting verification code: %s", err.Error())
	}

	cacheUser, err := s.cache.Get(ctx, "temp_user_"+hashEmail).Result()
	if err != nil {
		s.logger.Errorf("Error during getting temp user data: %s", err.Error())
		return errors.New(422, err.Error())
	}

	var tempUser dto.Register
	if err := json.Unmarshal([]byte(cacheUser), &tempUser); err != nil {
		s.logger.Errorf("Error during unmarshaling JSON to dto.Register: %s", err.Error())
		return err
	}

	var user model.User
	if err := utils.DtoToModel(&tempUser, &user); err != nil {
		s.logger.Errorf("Error during conversion: %s", err.Error())
		return err
	}

	if err := s.userService.Create(ctx, &user); err != nil {
		return err
	}

	if err := s.cache.Del(ctx, "temp_user_"+data.Email).Err(); err != nil {
		s.logger.Errorf("Error during deleting temp user data: %s", err.Error())
	}

	return nil
}
