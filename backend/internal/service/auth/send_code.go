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
	if err := mailer.SendMail(s.config.MailFrom, s.config.MailPassword, s.config.MailHost, s.config.MailPort, email); err != nil {
		return err
	}
	s.logger.Info("Email sent")
	s.logger.Info(code)
	return nil
}
