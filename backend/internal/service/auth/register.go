package auth

import (
	"context"
	"regexp"

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

	if err := s.userService.Create(ctx, &user); err != nil {
		return err
	}

	s.SendCode(ctx, dto.Email)

	return nil
}
