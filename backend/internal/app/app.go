package app

import (
	"fmt"
	"net/http"

	"github.com/Fi44er/sdmedik/backend/internal/config"
	events "github.com/Fi44er/sdmedik/backend/pkg/evenbus"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/Fi44er/sdmedik/backend/pkg/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger" // swagger handler
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron"
	"gorm.io/gorm"
)

type App struct {
	app             *fiber.App
	serviceProvider *serviceProvider
	httpService     *http.Server

	logger    *logger.Logger
	db        *gorm.DB
	validator *validator.Validate
	config    *config.Config
	cache     *redis.Client
	cron      *cron.Cron
}

func NewApp(logger *logger.Logger, db *gorm.DB, vavalidator *validator.Validate, config *config.Config, cache *redis.Client) (*App, error) {
	a := &App{
		app:       fiber.New(),
		logger:    logger,
		db:        db,
		validator: vavalidator,
		config:    config,
		cache:     cache,
		cron:      cron.New(),
	}

	return a, nil
}

func (a *App) Run() error {
	corsOrigins := fmt.Sprintf("http://127.0.0.1:8080, http://localhost:5173, %s", a.config.CorsOrigin)
	a.app.Use(cors.New(cors.Config{
		AllowOrigins:     corsOrigins, // Укажите источник вашего клиента
		AllowCredentials: true,        // Включение поддержки учетных данных
	}))

	if err := a.initDeps(); err != nil {
		return err
	}

	// a.cron.Start()
	// defer a.cron.Stop()
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
	eventBus := events.NewEventBus()
	a.serviceProvider, err = newServiceProvider(a.logger, a.db, a.validator, a.config, a.cache, a.cron, eventBus)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initRouter() error {
	a.app.Static("/api/v1/image", "./image")

	a.app.Get("/swagger/*", swagger.HandlerDefault)

	v1 := a.app.Group("/api/v1")

	user := v1.Group("/user")
	user.Get("/me", middleware.DeserializeUser(a.cache, a.db, a.config), a.serviceProvider.userProvider.UserImpl().GetMy)
	user.Get("/", a.serviceProvider.userProvider.UserImpl().GetAll)
	user.Get("/:id", a.serviceProvider.userProvider.UserImpl().GetByID)
	user.Put("/:id", a.serviceProvider.userProvider.UserImpl().Update)

	auth := v1.Group("/auth")
	auth.Post("/register", a.serviceProvider.authProvider.AuthImpl().Register)
	auth.Post("/login", a.serviceProvider.authProvider.AuthImpl().Login)
	auth.Post("/logout", middleware.DeserializeUser(a.cache, a.db, a.config), a.serviceProvider.authProvider.AuthImpl().Logout)
	auth.Post("/send-code", a.serviceProvider.authProvider.AuthImpl().SendCode)
	auth.Post("/verify-code", a.serviceProvider.authProvider.AuthImpl().VerifyCode)
	auth.Post("/refresh", a.serviceProvider.authProvider.AuthImpl().RefreshAccessToken)

	product := v1.Group("/product")
	product.Get("/filter/:category_id", a.serviceProvider.productProvider.ProductImpl().GetFilter)
	product.Get("/", a.serviceProvider.productProvider.ProductImpl().Get)
	product.Post("/", middleware.DeserializeUser(a.cache, a.db, a.config), a.serviceProvider.productProvider.ProductImpl().Create)
	product.Put("/:id", a.serviceProvider.productProvider.ProductImpl().Update)
	product.Delete("/:id", a.serviceProvider.productProvider.ProductImpl().Delete)
	product.Get("/top/:limit", a.serviceProvider.productProvider.ProductImpl().GetTopProducts)

	category := v1.Group("/category")
	category.Get("/", a.serviceProvider.categoryProvider.CategoryImpl().GetAll)
	category.Post("/", middleware.DeserializeUser(a.cache, a.db, a.config), a.serviceProvider.categoryProvider.CategoryImpl().Create)
	category.Get("/:id", a.serviceProvider.categoryProvider.CategoryImpl().GetByID)
	category.Delete("/:id", a.serviceProvider.categoryProvider.CategoryImpl().Delete)

	search := v1.Group("/search")
	search.Get("/", a.serviceProvider.searchProvider.SearchImpl().Search)

	basket := v1.Group("/basket")
	basket.Post("/create", a.serviceProvider.basketProvider.BasketImpl().Create)
	basket.Post("/", middleware.DeserializeUser(a.cache, a.db, a.config), a.serviceProvider.basketProvider.BasketImpl().AddItem)
	basket.Delete("/:id", middleware.DeserializeUser(a.cache, a.db, a.config), a.serviceProvider.basketProvider.BasketImpl().DeleteItem)
	basket.Get("/", middleware.DeserializeUser(a.cache, a.db, a.config), a.serviceProvider.basketProvider.BasketImpl().Get)

	webscraper := v1.Group("/webscraper")
	webscraper.Get("/", a.serviceProvider.webScraperProvider.WebScraperImpl().Scraper)

	order := v1.Group("/order")
	order.Post("/:id", a.serviceProvider.orderProvider.OrderImpl().NotAuthCreate)
	order.Post("/", middleware.DeserializeUser(a.cache, a.db, a.config), a.serviceProvider.orderProvider.OrderImpl().Create)
	order.Get("/my", middleware.DeserializeUser(a.cache, a.db, a.config), a.serviceProvider.orderProvider.OrderImpl().GetMyOrders)
	order.Get("/", a.serviceProvider.orderProvider.OrderImpl().GetAll)
	order.Put("/status", a.serviceProvider.orderProvider.OrderImpl().UpdateStatus)

	promotion := v1.Group("/promotion")
	promotion.Post("/", a.serviceProvider.promotionProvider.PromotionImpl().Create)
	promotion.Get("/", a.serviceProvider.promotionProvider.PromotionImpl().GetAll)
	promotion.Delete("/:id", a.serviceProvider.promotionProvider.PromotionImpl().Delete)

	return nil
}

func (a *App) runHttpServer() error {
	a.logger.Infof("HTTP server is running on %s", a.serviceProvider.HTTPConfig().Address())
	if err := a.app.Listen(a.serviceProvider.httpConfig.Address()); err != nil {
		a.logger.Fatal(err)
	}
	return nil
}
