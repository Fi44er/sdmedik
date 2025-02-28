package service

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/config"
	"github.com/Fi44er/sdmedik/backend/module/auth/dto"
	user_service "github.com/Fi44er/sdmedik/backend/module/user/service"
	"github.com/Fi44er/sdmedik/backend/shared/logger"
	"github.com/Fi44er/sdmedik/backend/shared/mailer"
	"github.com/redis/go-redis/v9"
)

var _ IAuthService = (*AuthService)(nil)

type IAuthService interface {
	Login(ctx context.Context, data *dto.LoginDTO) (*dto.LoginResponse, error)
	VerifyCode(ctx context.Context, data *dto.VerifyCodeDTO) error
	Register(ctx context.Context, data *dto.RegisterDTO) error
	SendCode(ctx context.Context, email string) error
	Logout(ctx context.Context, data *dto.LogoutDTO) error
	RefreshAccessToken(ctx context.Context, data *dto.RefreshTokenDTO) (string, error)
}

type AuthService struct {
	logger *logger.Logger
	cache  *redis.Client
	config *config.Config
	mailer *mailer.Mailer

	userServ user_service.IUserService
}

func NewAuthService(
	logger *logger.Logger,
	cache *redis.Client,
	config *config.Config,
	userServ user_service.IUserService,
) *AuthService {
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

	return &AuthService{
		logger:   logger,
		config:   config,
		cache:    cache,
		mailer:   m,
		userServ: userServ,
	}
}
