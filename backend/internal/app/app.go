package app

import (
	"net/http"

	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type App struct {
	app             *fiber.App
	serviceProvider *serviceProvider
	httpService     *http.Server
	logger          *logger.Logger
	db              *gorm.DB
}

func NewApp(logger *logger.Logger, db *gorm.DB) (*App, error) {
	a := &App{
		app:    fiber.New(),
		logger: logger,
		db:     db,
	}

	if err := a.initDeps(); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) Run() error {
	return a.runHttpServer()
}

func (a *App) initDeps() error {
	inits := []func() error{
		a.initConfig,
		a.initServiceProvider,
		a.initRouter,
	}

	for _, init := range inits {
		err := init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig() error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider() error {
	var err error
	a.serviceProvider, err = newServiceProvider(a.logger, a.db)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initRouter() error {
	v1 := a.app.Group("/api/v1")

	v1.Get("/hello", a.serviceProvider.userProvider.UserImpl().Hello)

	product := v1.Group("/product")
	product.Get("/", a.serviceProvider.productProvider.ProductImpl().GetAll)
	product.Post("/", a.serviceProvider.productProvider.ProductImpl().Create)
	return nil
}

func (a *App) runHttpServer() error {
	a.logger.Infof("HTTP server is running on %s", a.serviceProvider.HTTPConfig().Address())
	if err := a.app.Listen(a.serviceProvider.httpConfig.Address()); err != nil {
		a.logger.Fatal(err)
	}
	return nil
}
