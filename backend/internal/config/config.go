package config

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	HTTPHost            string        `mapstructure:"HTTP_HOST"`
	HTTPPort            string        `mapstructure:"HTTP_PORT"`
	PostgresUrl         string        `mapstructure:"POSTGRES_URL"`
	RedisUrl            string        `mapstructure:"REDIS_URL"`
	VerifyCodeExpiredIn time.Duration `mapstructure:"VERIFY_CODE_EXPIRED_IN"`
	ImageDir            string        `mapstructure:"IMAGE_DIR"`

	CorsOrigin string `mapstructure:"CORS_ORIGIN"`

	PayKeeperUser   string `mapstructure:"PAY_KEEPER_USER"`
	PayKeeperPass   string `mapstructure:"PAY_KEEPER_PASS"`
	PayKeeperServer string `mapstructure:"PAY_KEEPER_SERVER"`

	MailTemplatePath string `mapstructure:"MAIL_TEMPLATE_PATH"`
	MailHost         string `mapstructure:"MAIL_HOST"`
	MailUser         string `mapstructure:"MAIL_USER"`
	MailPassword     string `mapstructure:"MAIL_PASSWORD"`
	MailFrom         string `mapstructure:"MAIL_FROM"`
	MailPort         string `mapstructure:"MAIL_PORT"`

	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAX_AGE"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAX_AGE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}
