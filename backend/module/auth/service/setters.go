package service

import (
	"context"
	"encoding/json"
	"regexp"
	"time"

	"github.com/Fi44er/sdmedik/backend/converter"
	"github.com/Fi44er/sdmedik/backend/module/auth/dto"
	customerr "github.com/Fi44er/sdmedik/backend/shared/custom_err"
	"github.com/Fi44er/sdmedik/backend/shared/utils"
	"github.com/redis/go-redis/v9"
)

const (
	CodeRedisPrefix = "verification_codes_"
	UserRedisPrefix = "temp_user_"
)

func (s *AuthService) getFromCache(ctx context.Context, key string) (string, error) {
	val, err := s.cache.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", customerr.NewError(404, "Key not found")
	} else if err != nil {
		return "", customerr.NewError(500, err.Error())
	}
	return val, nil
}

func (s *AuthService) setToCache(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	if err := s.cache.Set(ctx, key, value, duration).Err(); err != nil {
		s.logger.Errorf("Error setting cache key %s: %s", key, err.Error())
		return customerr.NewError(500, err.Error())
	}
	return nil
}

func (s *AuthService) deleteFromCache(ctx context.Context, key string) error {
	if err := s.cache.Del(ctx, key).Err(); err != nil {
		s.logger.Errorf("Error deleting cache key %s: %s", key, err.Error())
		return customerr.NewError(500, err.Error())
	}
	return nil
}

func (s *AuthService) createAndStoreToken(ctx context.Context, userID string, expiresIn time.Duration, privateKey string, userAgent string) (string, error) {
	tokenDetails, err := utils.CreateToken(userID, expiresIn, privateKey)
	if err != nil {
		return "", customerr.NewError(422, err.Error())
	}
	key := userAgent + ":" + tokenDetails.TokenUUID
	err = s.setToCache(ctx, key, userID, time.Until(time.Unix(*tokenDetails.ExpiresIn, 0)))
	return *tokenDetails.Token, err
}

func (s *AuthService) Login(ctx context.Context, data *dto.LoginDTO) (*dto.LoginResponse, error) {
	user, err := s.userServ.GetByEmail(ctx, data.Email)
	if err != nil || !utils.ComparePassword(user.Password, data.Password) {
		return nil, customerr.ErrInvalidEmailOrPassword
	}

	accessToken, err := s.createAndStoreToken(ctx, user.ID, s.config.AccessTokenExpiresIn, s.config.AccessTokenPrivateKey, data.UserAgent)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.createAndStoreToken(ctx, user.ID, s.config.RefreshTokenExpiresIn, s.config.RefreshTokenPrivateKey, data.UserAgent)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s *AuthService) VerifyCode(ctx context.Context, data *dto.VerifyCodeDTO) error {
	hashEmail, err := utils.HashString(data.Email)
	if err != nil {
		return customerr.NewError(500, err.Error())
	}

	code, err := s.getFromCache(ctx, CodeRedisPrefix+hashEmail)
	if err != nil || code != data.Code {
		return customerr.ErrInvalidCode
	}

	_ = s.deleteFromCache(ctx, CodeRedisPrefix+hashEmail)

	cacheUser, err := s.getFromCache(ctx, UserRedisPrefix+hashEmail)
	if err != nil {
		return err
	}

	var tempUser dto.RegisterDTO
	if err := json.Unmarshal([]byte(cacheUser), &tempUser); err != nil {
		return err
	}

	user := converter.ConvertRegisterDtoToUserDomain(&tempUser)
	if err := s.userServ.Create(ctx, user); err != nil {
		return err
	}

	return s.deleteFromCache(ctx, UserRedisPrefix+hashEmail)
}

func (s *AuthService) Register(ctx context.Context, data *dto.RegisterDTO) error {
	data.PhoneNumber = regexp.MustCompile("[^0-9]").ReplaceAllString(data.PhoneNumber, "")
	if len(data.PhoneNumber) != 11 {
		return customerr.ErrInvalidPhoneNumber
	}

	if user, _ := s.userServ.GetByEmail(ctx, data.Email); user != nil {
		return customerr.ErrUserAlreadyExists
	}

	data.Password = utils.GeneratePassword(data.Password)

	hashEmail, err := utils.HashString(data.Email)
	if err != nil {
		return err
	}

	dataCache, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := s.setToCache(ctx, UserRedisPrefix+hashEmail, dataCache, 10*time.Minute); err != nil {
		return err
	}

	return s.SendCode(ctx, data.Email)
}

func (s *AuthService) SendCode(ctx context.Context, email string) error {
	code := utils.GenerateCode(6)
	hashEmail, err := utils.HashString(email)
	if err != nil {
		return err
	}

	if _, err := s.getFromCache(ctx, UserRedisPrefix+hashEmail); err != nil {
		return customerr.NewError(422, err.Error())
	}

	if err := s.setToCache(ctx, CodeRedisPrefix+hashEmail, code, s.config.VerifyCodeExpiredIn); err != nil {
		return err
	}

	templateData := struct{ Code string }{Code: code}

	s.mailer.SendMailAsync(
		s.config.MailFrom,
		email,
		"Код подтверждения регистрации",
		templateData,
	)

	s.logger.Info(code)
	return nil
}

func (s *AuthService) RefreshAccessToken(ctx context.Context, data *dto.RefreshTokenDTO) (string, error) {
	if data.RefreshToken == "" {
		return "", customerr.ErrInvalidToken
	}

	tokenClaims, err := utils.ValidateToken(data.RefreshToken, s.config.RefreshTokenPublicKey)
	if err != nil {
		return "", customerr.NewError(403, err.Error())
	}

	userID, err := s.getFromCache(ctx, tokenClaims.TokenUUID)
	if err != nil {
		return "", customerr.ErrCouldNotRefreshToken
	}

	user, err := s.userServ.GetByID(ctx, userID)
	if err != nil || user == nil {
		return "", customerr.ErrUserNotFound
	}

	return s.createAndStoreToken(ctx, user.ID, s.config.AccessTokenExpiresIn, s.config.AccessTokenPrivateKey, data.UserAgent)
}

func (s *AuthService) Logout(ctx context.Context, data *dto.LogoutDTO) error {
	if data.RefreshToken == "" {
		return customerr.ErrInvalidToken
	}

	tokenClaims, err := utils.ValidateToken(data.RefreshToken, s.config.RefreshTokenPublicKey)
	if err != nil {
		return customerr.NewError(401, err.Error())
	}

	if err := s.deleteFromCache(ctx, tokenClaims.TokenUUID); err != nil {
		return err
	}

	return nil
}
