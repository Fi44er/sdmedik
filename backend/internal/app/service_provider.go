package app

import (
	"github.com/Fi44er/sdmedik/backend/internal/app/provider"
	"github.com/Fi44er/sdmedik/backend/internal/config"
	events "github.com/Fi44er/sdmedik/backend/pkg/evenbus"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron"
	"gorm.io/gorm"
)

type serviceProvider struct {
	httpConfig config.HTTPConfig

	userProvider                provider.UserProvider
	productProvider             provider.ProductProvider
	authProvider                provider.AuthProvider
	categoryProvider            provider.CategoryProvider
	characteristicProvider      provider.CharacteristicProvider
	transactionManagerProvider  provider.TransactionManagerProvider
	characteristicValueProvider provider.CharacteristicValueProvider
	imageProvider               provider.ImageProvider
	searchProvider              provider.SearchProvider
	indexProvider               provider.IndexProvider
	basketProvider              provider.BasketProvider
	webScraperProvider          provider.WebscraperProvider
	certificateProvider         provider.CertificateProvider
	orderProvider               provider.OrderProvider

	logger    *logger.Logger
	db        *gorm.DB
	validator *validator.Validate
	config    *config.Config
	cache     *redis.Client
	cron      *cron.Cron

	eventBus *events.EventBus
}

func newServiceProvider(
	logger *logger.Logger,
	db *gorm.DB,
	validator *validator.Validate,
	config *config.Config,
	cache *redis.Client,
	cron *cron.Cron,
	eventBus *events.EventBus,
) (*serviceProvider, error) {
	a := &serviceProvider{
		logger:    logger,
		db:        db,
		validator: validator,
		config:    config,
		cache:     cache,
		cron:      cron,
		eventBus:  eventBus,
	}

	if err := a.initDeps(); err != nil {
		return nil, err
	}

	return a, nil
}

func (s *serviceProvider) initDeps() error {
	inits := []func() error{
		s.initTransactionManagerProvider,
		s.initCharacteristicValueProvider,
		s.initCharacteristicProvider,
		s.initImageProvider,

		s.initCategoryProvider,
		s.initCertificateProvider,
		s.initProductProvider,
		s.initWebScraperProvider,
		s.initBasketProvider,
		s.initUserProvider,
		s.initAuthProvider,
		s.initIndexProvider,
		s.initSearchProvider,
		s.initOrderProvider,
	}

	for _, init := range inits {
		err := init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *serviceProvider) initTransactionManagerProvider() error {
	s.transactionManagerProvider = *provider.NewTransactionManagerProvider(s.logger, s.db)
	return nil
}

func (s *serviceProvider) initUserProvider() error {
	s.userProvider = *provider.NewUserProvider(s.logger, s.validator, s.db, s.config, s.cache, s.basketProvider.BasketService())
	return nil
}

func (s *serviceProvider) initAuthProvider() error {
	s.authProvider = *provider.NewAuthProvider(s.logger, s.validator, s.config, s.cache, s.userProvider.UserService())
	return nil
}

func (s *serviceProvider) initImageProvider() error {
	s.imageProvider = *provider.NewImageProvider(s.logger, s.db, s.validator, s.config)
	return nil
}

func (s *serviceProvider) initProductProvider() error {
	s.productProvider = *provider.NewProductProvider(
		s.logger,
		s.db,
		s.validator,
		s.eventBus,
		s.cache,
		s.categoryProvider.CategoryService(),
		s.characteristicValueProvider.CharacteristicValueService(),
		s.transactionManagerProvider.TransactionManager(),
		s.imageProvider.ImageService(),
		s.characteristicProvider.CharacteristicService(),
		s.certificateProvider.CertificateService(),
	)
	return nil
}

func (s *serviceProvider) initCharacteristicProvider() error {
	s.characteristicProvider = *provider.NewCharacteristicProvider(s.logger, s.db, s.validator)
	return nil
}

func (s *serviceProvider) initCharacteristicValueProvider() error {
	s.characteristicValueProvider = *provider.NewChracteristicValueProvider(s.logger, s.db, s.validator)
	return nil
}

func (s *serviceProvider) initCategoryProvider() error {
	s.categoryProvider = *provider.NewCategoryProvider(
		s.logger,
		s.db, s.validator,
		s.characteristicProvider.CharacteristicService(),
		s.transactionManagerProvider.TransactionManager(),
		s.imageProvider.ImageService(),
		s.eventBus,
	)
	return nil
}

func (s *serviceProvider) initIndexProvider() error {
	s.indexProvider = *provider.NewIndexProvider(
		s.logger,
		s.validator,
		s.productProvider.ProductService(),
		s.categoryProvider.CategoryService(),
		s.eventBus,
	)
	return nil
}

func (s *serviceProvider) initSearchProvider() error {
	s.searchProvider = *provider.NewSearchProvider(s.logger, s.validator, s.indexProvider.IndexService())
	return nil
}

func (s *serviceProvider) initBasketProvider() error {
	s.basketProvider = *provider.NewBasketProvider(s.logger, s.db, s.validator, s.productProvider.ProductService())
	return nil
}

func (s *serviceProvider) initWebScraperProvider() error {
	s.webScraperProvider = *provider.NewWebscraperProvider(s.logger, s.validator, s.cron, s.certificateProvider.CertificateService(), s.productProvider.ProductService())
	s.webScraperProvider.WebScraperService()
	return nil
}

func (s *serviceProvider) initCertificateProvider() error {
	s.certificateProvider = *provider.NewCertificateProvider(s.logger, s.db)
	return nil
}

func (s *serviceProvider) initOrderProvider() error {
	s.orderProvider = *provider.NewOrderProvider(s.logger, s.validator, s.db, s.config, s.basketProvider.BasketService(), s.certificateProvider.CertificateService())
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
