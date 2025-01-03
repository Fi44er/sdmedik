package auth

import (
	"context"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) SendCode(ctx context.Context, email string) error {
	code := utils.GenerateCode(6)
	hashEmail, err := utils.HashString(email)
	if err != nil {
		s.logger.Errorf("Error during generating hash: %s", err.Error())
		return err
	}

	expiredIn := s.config.VerifyCodeExpiredIn

	_, err = s.cache.Get(ctx, "temp_user_"+hashEmail).Result()
	if err != nil {
		s.logger.Errorf("Error during getting temp user data: %s", err.Error())
		return errors.New(422, err.Error())
	}

	if err := s.cache.Set(ctx, "verification_codes_"+hashEmail, code, expiredIn).Err(); err != nil {
		s.logger.Errorf("Error during saving verification code: %s", err.Error())
		return errors.New(422, err.Error())
	}

	templateData := struct {
		Code string
	}{
		Code: code,
	}

	s.mailer.SendMailAsync(
		s.config.MailFrom, // Отправитель
		email,             // Получатель
		"Код подтверждения регистрации", // Тема письма
		templateData, // Данные для шаблона
	)
	s.logger.Info(code)
	return nil
}
