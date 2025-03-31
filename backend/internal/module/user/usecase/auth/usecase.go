package auth

import (
	"context"
	"regexp"
	"time"

	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/internal/module/user/dto"
	"github.com/Fi44er/sdmedik/backend/internal/module/user/entity"
	"github.com/Fi44er/sdmedik/backend/internal/module/user/pkg/constant"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/Fi44er/sdmedik/backend/pkg/mailer"
	"github.com/Fi44er/sdmedik/backend/pkg/redis"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

type IUserUsecase interface {
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByID(ctx context.Context, id string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
}

type AuthUsecase struct {
	logger       *logger.Logger
	redisManager redis.IRedisManager
	config       *config.Config
	mailer       *mailer.Mailer
	userUsecase  IUserUsecase
}

func NewAuthUsecase(
	logger *logger.Logger,
	redisManager redis.IRedisManager,
	config *config.Config,
	userUsecase IUserUsecase,
) *AuthUsecase {
	templatePath := config.MailTemplatePath
	m, err := mailer.NewMailer(
		config.MailHost,           // SMTP-хост
		config.MailPort,           // Порт
		config.MailFrom,           // Ваш email
		config.MailPassword,       // Пароль от почты
		templatePath+"index.html", // Путь к шаблону
		5,                         // Размер пула соединений
	)

	if err != nil {
		logger.Fatalf("Failed to initialize mailer: %v", err)
		return nil
	}

	return &AuthUsecase{
		logger:       logger,
		config:       config,
		redisManager: redisManager,
		mailer:       m,
		userUsecase:  userUsecase,
	}
}

const (
	CodeRedisPrefix = "verification_codes_"
	UserRedisPrefix = "temp_user_"
)

func (s *AuthUsecase) createAndStoreToken(ctx context.Context, userID string, expiresIn time.Duration, privateKey string, userAgent string) (string, error) {
	tokenDetails, err := utils.CreateToken(userID, expiresIn, privateKey)
	if err != nil {
		return "", constant.ErrUnprocessableEntity
	}
	key := userAgent + ":" + tokenDetails.TokenUUID
	err = s.redisManager.Set(ctx, key, userID, time.Until(time.Unix(*tokenDetails.ExpiresIn, 0)))
	return *tokenDetails.Token, err
}

func (s *AuthUsecase) SignIn(ctx context.Context, data *dto.LoginDTO) (*dto.LoginResponse, error) {
	user, err := s.userUsecase.GetByEmail(ctx, data.Email)
	if err != nil || !utils.ComparePassword(user.PasswordHash, data.Password) {
		return nil, constant.ErrInvalidEmailOrPassword
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

func (s *AuthUsecase) VerifyCode(ctx context.Context, data *dto.VerifyCodeDTO) error {
	hashEmail, err := utils.HashString(data.Email)
	if err != nil {
		return constant.ErrInternalServerError
	}

	var code string
	if err := s.redisManager.Get(ctx, CodeRedisPrefix+hashEmail, &code); err != nil {
		return constant.ErrInternalServerError
	}

	if err := s.redisManager.Del(ctx, CodeRedisPrefix+hashEmail); err != nil {
		return err
	}

	var tempUser dto.RegisterDTO
	if err := s.redisManager.Get(ctx, UserRedisPrefix+hashEmail, &tempUser); err != nil {
		return err
	}

	// user := converter.ConvertRegisterDtoToUserDomain(&tempUser)
	// if err := s.userUsecase.Create(ctx, user); err != nil {
	// 	return err
	// }

	return s.redisManager.Del(ctx, UserRedisPrefix+hashEmail)
}

func (s *AuthUsecase) SignUp(ctx context.Context, data *dto.RegisterDTO) error {
	data.PhoneNumber = regexp.MustCompile("[^0-9]").ReplaceAllString(data.PhoneNumber, "")
	if len(data.PhoneNumber) != 11 {
		return constant.ErrInvalidPhoneNumber
	}

	user, err := s.userUsecase.GetByEmail(ctx, data.Email)
	if err != nil || user != nil {
		return constant.ErrUserAlreadyExists
	}

	data.Password = utils.GeneratePassword(data.Password)

	hashEmail, err := utils.HashString(data.Email)
	if err != nil {
		return err
	}

	if err := s.redisManager.Set(ctx, UserRedisPrefix+hashEmail, data, 10*time.Minute); err != nil {
		return err
	}

	return s.SendCode(ctx, data.Email)
}

func (s *AuthUsecase) SendCode(ctx context.Context, email string) error {
	code := utils.GenerateCode(6)
	hashEmail, err := utils.HashString(email)
	if err != nil {
		return err
	}

	var tempUser dto.RegisterDTO
	if err := s.redisManager.Get(ctx, UserRedisPrefix+hashEmail, &tempUser); err != nil {
		return constant.ErrUnprocessableEntity
	}

	if err := s.redisManager.Set(ctx, CodeRedisPrefix+hashEmail, code, s.config.VerifyCodeExpiredIn); err != nil {
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

func (s *AuthUsecase) RefreshAccessToken(ctx context.Context, data *dto.RefreshTokenDTO) (string, error) {
	if data.RefreshToken == "" {
		return "", constant.ErrInvalidToken
	}

	tokenClaims, err := utils.ValidateToken(data.RefreshToken, s.config.RefreshTokenPublicKey)
	if err != nil {
		return "", constant.ErrForbidden
	}

	var userID string
	if err := s.redisManager.Get(ctx, tokenClaims.TokenUUID, &userID); err != nil {
		return "", constant.ErrCouldNotRefreshToken
	}

	user, err := s.userUsecase.GetByID(ctx, userID)
	if err != nil || user == nil {
		return "", constant.ErrUserNotFound
	}

	return s.createAndStoreToken(ctx, user.ID, s.config.AccessTokenExpiresIn, s.config.AccessTokenPrivateKey, data.UserAgent)
}

func (s *AuthUsecase) SignOut(ctx context.Context, data *dto.LogoutDTO) error {
	if data.RefreshToken == "" {
		return constant.ErrInvalidToken
	}

	tokenClaims, err := utils.ValidateToken(data.RefreshToken, s.config.RefreshTokenPublicKey)
	if err != nil {
		return constant.ErrAnauthorized
	}

	if err := s.redisManager.Del(ctx, tokenClaims.TokenUUID); err != nil {
		return err
	}

	return nil
}
