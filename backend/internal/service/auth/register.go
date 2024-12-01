package auth

import (
	"context"
	"encoding/json"
	"regexp"
	"time"

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

	re := regexp.MustCompile("[^0-9]")
	dto.PhoneNumber = re.ReplaceAllString(dto.PhoneNumber, "")
	if len(dto.PhoneNumber) != 11 {
		return errors.New(400, "Invalid phone number")
	}

	var user model.User
	existUser, err := s.userService.GetByEmail(ctx, dto.Email)
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

	s.SendCode(ctx, dto.Email)

	hashEmail, err := utils.HashString(dto.Email)
	if err != nil {
		s.logger.Errorf("Error during generating hash: %s", err.Error())
		return err
	}

	data, err := json.Marshal(dto)
	if err != nil {
		s.logger.Errorf("Error during marshaling dto to JSON: %s", err.Error())
		return err
	}

	if err := s.cache.Set(ctx, "temp_user_"+hashEmail, data, time.Minute*10).Err(); err != nil {
		s.logger.Errorf("Error during caching temp user data: %s", err.Error())
		return err
	}

	return nil
}
