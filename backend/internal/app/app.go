package app

import (
	"log"
	"net/http"

	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	app            *fiber.App
	sericeProvider *serviceProvider
	httpService    *http.Server
}

func NewApp() (*App, error) {
	a := &App{app: fiber.New()}

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
	a.sericeProvider = newServiceProvider()
	return nil
}

func (a *App) initRouter() error {
	v1 := a.app.Group("/api/v1")

	v1.Get("/hello", a.sericeProvider.userProvider.UserImpl().Hello)
	return nil
}

func (a *App) runHttpServer() error {
	log.Printf("HTTP server is running on %s", a.sericeProvider.HTTPConfig().Address())
	if err := a.initRouter(); err != nil {
		log.Fatal(err)
	}
	if err := a.app.Listen(a.sericeProvider.httpConfig.Address()); err != nil {
		log.Fatal(err)
	}
	return nil
}
