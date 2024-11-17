package app

import (
	"github.com/Fi44er/sdmedik/backend/internal/app/provider"
	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

type serviceProvider struct {
	httpConfig      config.HTTPConfig
	userProvider    provider.UserProvider
	productProvider provider.ProductProvider

	logger *logger.Logger
	db     *gorm.DB
}

func newServiceProvider(logger *logger.Logger, db *gorm.DB) (*serviceProvider, error) {
	a := &serviceProvider{
		logger: logger,
		db:     db,
	}

	if err := a.initDeps(); err != nil {
		return nil, err
	}

	return a, nil
}

func (s *serviceProvider) initDeps() error {
	inits := []func() error{
		s.initUserProvider,
		s.initProductProvider,
	}

	for _, init := range inits {
		err := init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *serviceProvider) initUserProvider() error {
	s.userProvider = *provider.NewUserProvider(s.logger)
	return nil
}

func (s *serviceProvider) initProductProvider() error {
	s.productProvider = *provider.NewProductProvider(s.logger, s.db)
	return nil
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			s.logger.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}
