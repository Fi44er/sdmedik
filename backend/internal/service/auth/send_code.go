package auth

import (
	"context"
	"encoding/hex"

	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/mailer"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
	"golang.org/x/crypto/blake2b"
)

func (s *service) SendCode(ctx context.Context, email string) error {
	code := utils.GenerateCode(6)
	h, err := blake2b.New256(nil)
	if err != nil {
		return err
	}
	h.Write([]byte(email))
	hashBytes := h.Sum(nil)
	hashEmail := hex.EncodeToString(hashBytes)

	expiredIn := s.config.VerifyCodeExpiredIn
	if err := s.cache.Set(ctx, "verification_codes_"+hashEmail, code, expiredIn).Err(); err != nil {
		return errors.New(422, err.Error())
	}

	mailer.SendMail(s.config.MailFrom, s.config.MailPassword, s.config.MailHost, s.config.MailPort, email)

	return nil
}
