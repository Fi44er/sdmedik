package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

const (
	RESET_PASS_CODE = "reset_pass_code_"
)

func (s *service) ResetPassword(ctx context.Context, email string) error {
	user, err := s.userService.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New(400, "User not found")
	}

	s.logger.Infof("User: %v", user)

	token, err := GenerateSecureToken(32)
	if err != nil {
		return err
	}

	if err := s.cache.Set(ctx, RESET_PASS_CODE+token, user.ID, s.config.RessetPassCodeExpiredIn).Err(); err != nil {
		return err
	}

	resetLink := fmt.Sprintf(s.config.FrontendURL+s.config.ResetPassLink+"?token=%s", token)
	templateData := struct {
		Code string
	}{
		Code: resetLink,
	}

	s.mailer.SendMailAsync(
		s.config.MailFrom, // Отправитель
		email,             // Получатель
		"Код подтверждения регистрации", // Тема письма
		templateData, // Данные для шаблона
	)

	return nil
}

func (s *service) ChangePassword(ctx context.Context, token, password, userID string) error {
	updateUserData := dto.UpdateUser{
		Password: utils.GeneratePassword(password),
	}

	if err := s.userService.Update(ctx, &updateUserData, userID); err != nil {
		return err
	}

	if err := s.cache.Del(ctx, RESET_PASS_CODE+token).Err(); err != nil {
		return err
	}

	return nil
}

func (s *service) ValidateToken(ctx context.Context, token string) (string, error) {
	userID, err := s.cache.Get(ctx, RESET_PASS_CODE+token).Result()
	s.logger.Infof("userID: %s", token)
	if err != nil || userID == "" {
		return "", errors.New(400, "Invalid token")
	}

	return userID, nil
}

func GenerateSecureToken(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
