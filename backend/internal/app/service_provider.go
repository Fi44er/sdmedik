package app

import (
	"log"

	"github.com/Fi44er/sdmedik/backend/internal/app/provider"
	"github.com/Fi44er/sdmedik/backend/internal/config"
)

type serviceProvider struct {
	httpConfig   config.HTTPConfig
	userProvider provider.UserProvider
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}
