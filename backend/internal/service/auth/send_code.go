package auth

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/mailer"
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
	if err := s.cache.Set(ctx, "verification_codes_"+hashEmail, code, expiredIn).Err(); err != nil {
		s.logger.Errorf("Error during saving verification code: %s", err.Error())
		return errors.New(422, err.Error())
	}

	s.logger.Info("Sending email verification code...")
	m, err := mailer.NewMailer(
		s.config.MailHost,                // SMTP-хост
		s.config.MailPort,                // Порт
		s.config.MailFrom,                // Ваш email
		s.config.MailPassword,            // Пароль от почты
		"pkg/mailer/template/index.html", // Путь к шаблону
		5,                                // Размер пула соединений
	)

	if err != nil {
		s.logger.Fatalf("Failed to initialize mailer: %v", err)
		return err
	}

	templateData := struct {
		Code string
	}{
		Code: code,
	}

	m.SendMailAsync(
		s.config.MailFrom, // Отправитель
		email,             // Получатель
		"Код подтверждения регистрации", // Тема письма
		templateData, // Данные для шаблона
	)
	s.logger.Info(code)
	return nil
}
