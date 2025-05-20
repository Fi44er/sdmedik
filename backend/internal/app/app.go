package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Fi44er/sdmedik/backend/internal/config"
	events "github.com/Fi44er/sdmedik/backend/pkg/evenbus"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/Fi44er/sdmedik/backend/pkg/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	redis_store "github.com/gofiber/storage/redis"
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
	sessStore *session.Store
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
	corsOrigins := fmt.Sprintf("http://127.0.0.1:8080, http://localhost:5173, http://localhost:8080, %s", a.config.CorsOrigin)
	a.app.Use(cors.New(cors.Config{
		AllowOrigins:     corsOrigins, // Укажите источник вашего клиента
		AllowCredentials: true,        // Включение поддержки учетных данных
	}))

	a.app.Use(func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return c.Next()
	})

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
		a.initSessionStore,
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

func (a *App) initSessionStore() error {
	storage := redis_store.New(redis_store.Config{
		Host:     "127.0.0.1",
		Port:     6379,
		Password: "", // если есть пароль, укажите его здесь
		Database: 0,  // используйте 0 для дефолтной БД Redis
		Reset:    false,
	})

	store := session.New(session.Config{
		Storage:    storage,        // Подключаем Redis как хранилище
		Expiration: 24 * time.Hour, // Время жизни сессии
	})

	a.sessStore = store

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

	allowGuest := middleware.AllowGuest(a.cache, a.db, a.config, a.sessStore)
	deserializeUser := middleware.DeserializeUser(a.cache, a.db, a.config)
	adminRoleRequired := middleware.RoleRequired("admin")

	v1 := a.app.Group("/api/v1")

	user := v1.Group("/user")
	user.Get("/me", deserializeUser, a.serviceProvider.userProvider.UserImpl().GetMy)
	user.Get("/", a.serviceProvider.userProvider.UserImpl().GetAll)
	user.Get("/:id", a.serviceProvider.userProvider.UserImpl().GetByID)
	user.Put("/:id", a.serviceProvider.userProvider.UserImpl().Update)

	auth := v1.Group("/auth")
	auth.Post("/register", a.serviceProvider.authProvider.AuthImpl().Register)
	auth.Post("/login", allowGuest, a.serviceProvider.authProvider.AuthImpl().Login)
	auth.Post("/logout", deserializeUser, a.serviceProvider.authProvider.AuthImpl().Logout)
	auth.Post("/send-code", a.serviceProvider.authProvider.AuthImpl().SendCode)
	auth.Post("/verify-code", a.serviceProvider.authProvider.AuthImpl().VerifyCode)
	auth.Post("/refresh", a.serviceProvider.authProvider.AuthImpl().RefreshAccessToken)

	auth.Get("/reset-password/:email", a.serviceProvider.authProvider.AuthImpl().ResetPassword)
	auth.Post("/reset-password", a.serviceProvider.authProvider.AuthImpl().ChangePassword)

	product := v1.Group("/product")
	product.Get("/filter/:category_id", a.serviceProvider.productProvider.ProductImpl().GetFilter)
	product.Get("/", a.serviceProvider.productProvider.ProductImpl().Get)
	product.Post("/", deserializeUser, adminRoleRequired, a.serviceProvider.productProvider.ProductImpl().Create)
	product.Put("/:id", deserializeUser, adminRoleRequired, a.serviceProvider.productProvider.ProductImpl().Update)
	product.Delete("/:id", deserializeUser, adminRoleRequired, a.serviceProvider.productProvider.ProductImpl().Delete)
	product.Get("/top/:limit", a.serviceProvider.productProvider.ProductImpl().GetTopProducts)

	category := v1.Group("/category")
	category.Get("/", a.serviceProvider.categoryProvider.CategoryImpl().GetAll)
	category.Post("/", deserializeUser, adminRoleRequired, a.serviceProvider.categoryProvider.CategoryImpl().Create)
	category.Get("/:id", a.serviceProvider.categoryProvider.CategoryImpl().GetByID)
	category.Delete("/:id", deserializeUser, adminRoleRequired, a.serviceProvider.categoryProvider.CategoryImpl().Delete)
	category.Put("/:id", deserializeUser, adminRoleRequired, a.serviceProvider.categoryProvider.CategoryImpl().Update)

	search := v1.Group("/search")
	search.Get("/", a.serviceProvider.searchProvider.SearchImpl().Search)

	basket := v1.Group("/basket")
	basket.Post("/create", a.serviceProvider.basketProvider.BasketImpl().Create)
	basket.Post("/", allowGuest, a.serviceProvider.basketProvider.BasketImpl().AddItem)
	basket.Delete("/:id", allowGuest, a.serviceProvider.basketProvider.BasketImpl().DeleteItem)
	basket.Get("/", allowGuest, a.serviceProvider.basketProvider.BasketImpl().Get)

	webscraper := v1.Group("/webscraper")
	webscraper.Post("/start/", deserializeUser, adminRoleRequired, a.serviceProvider.webScraperProvider.WebScraperImpl().Scraper)
	webscraper.Post("/cancel/", deserializeUser, adminRoleRequired, a.serviceProvider.webScraperProvider.WebScraperImpl().CancelScraper)

	order := v1.Group("/order")
	order.Post("/:id", a.serviceProvider.orderProvider.OrderImpl().NotAuthCreate)
	order.Post("/", allowGuest, a.serviceProvider.orderProvider.OrderImpl().Create)
	order.Get("/my", deserializeUser, a.serviceProvider.orderProvider.OrderImpl().GetMyOrders)
	order.Get("/", deserializeUser, adminRoleRequired, a.serviceProvider.orderProvider.OrderImpl().GetAll)
	order.Put("/status", deserializeUser, adminRoleRequired, a.serviceProvider.orderProvider.OrderImpl().UpdateStatus)

	promotion := v1.Group("/promotion")
	promotion.Post("/", deserializeUser, adminRoleRequired, a.serviceProvider.promotionProvider.PromotionImpl().Create)
	promotion.Get("/", a.serviceProvider.promotionProvider.PromotionImpl().GetAll)
	promotion.Delete("/:id", deserializeUser, adminRoleRequired, a.serviceProvider.promotionProvider.PromotionImpl().Delete)

	chat := v1.Group("/chat")
	chat.Get("/conn/:user_id", socketio.New(a.serviceProvider.chatProvider.ChatImpl().WS()))

	return nil
}

func (a *App) runHttpServer() error {
	a.logger.Infof("HTTP server is running on %s", a.serviceProvider.HTTPConfig().Address())
	if err := a.app.Listen(a.serviceProvider.httpConfig.Address()); err != nil {
		a.logger.Fatal(err)
	}
	return nil
}
