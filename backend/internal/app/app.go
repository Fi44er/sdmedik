package app

import (
	"flag"
	"fmt"
	"sync"

	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/Fi44er/sdmedik/backend/pkg/middleware"
	"github.com/Fi44er/sdmedik/backend/pkg/postgres"
	redis_connect "github.com/Fi44er/sdmedik/backend/pkg/redis"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	app *fiber.App

	config     *config.Config
	logger     *logger.Logger
	validator  *validator.Validate
	httpConfig config.HTTPConfig

	db    *gorm.DB
	redis *redis.Client

	moduleProvider *moduleProvider

	migrate   bool // Флаг для миграции
	redisMode int  // Флаг для режима Redis
}

func NewApp() *App {
	migrate := flag.Bool("migrate", false, "Run database migration on startup (true/false)")
	redisMode := flag.Int("redis", 0, "Redis cache mode: 0 - no flush, 1 - selective flush, 2 - full flush")
	flag.Parse()

	return &App{
		app:       fiber.New(),
		migrate:   *migrate,
		redisMode: *redisMode,
	}
}

var wg sync.WaitGroup

func (app *App) Run() error {
	app.app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://127.0.0.1:8080, http://localhost:5173, http://localhost:8080",
		AllowCredentials: true,
	}))

	app.app.Use(logger.LoggerMiddleware())
	app.app.Use(middleware.ErrHandler)

	err := app.initDeps()
	if err != nil {
		return err
	}

	return app.runHttpServer()
}

func (app *App) initDeps() error {
	inits := []func() error{
		app.initConfig,
		app.initLogger,
		app.initValidator,
		app.initDb,
		app.initRedis,
		app.initModuleProvider,
		app.initRouter,
	}
	for _, init := range inits {
		err := init()
		if err != nil {
			return fmt.Errorf("✖ Failed to initialize dependencies: %s", err.Error())
		}
	}
	return nil
}

func (app *App) initConfig() error {
	if app.config == nil {
		config, err := config.LoadConfig(".")
		if err != nil {
			return fmt.Errorf("✖ Failed to load config: %s", err.Error())
		}
		app.config = config
	}

	err := config.Load(".env")
	if err != nil {
		return fmt.Errorf("✖ Failed to load config: %s", err.Error())
	}

	return nil
}

func (app *App) initDb() error {
	if app.db == nil {
		db, err := postgres.ConnectDb(app.config.PostgresUrl, app.logger)
		if err != nil {
			return err
		}
		app.db = db

		// Используем значение migrate из структуры App
		if err := postgres.Migrate(db, app.migrate, app.logger); err != nil {
			return fmt.Errorf("✖ Failed to migrate database: %s", err.Error())
		}
	}

	return nil
}

func (app *App) initRedis() error {
	if app.redis == nil {
		redis, err := redis_connect.Connect(app.config.RedisUrl, app.logger)
		if err != nil {
			app.logger.Errorf("Failed to connect to Redis: %v", err)
			return nil
		}
		app.redis = redis

		// Используем значение redisMode из структуры App
		if err := redis_connect.FlushRedisCache(redis, app.redisMode, app.logger); err != nil {
			err = fmt.Errorf("✖ Failed to flush redis cache: %v", err)
			app.logger.Errorf("%s", err.Error())
			return err
		}
	}
	return nil
}

// Остальные методы остаются без изменений
func (app *App) initLogger() error {
	if app.logger == nil {
		app.logger = logger.NewLogger()
	}
	return nil
}

func (app *App) initValidator() error {
	if app.validator == nil {
		app.validator = validator.New()
	}
	return nil
}

func (app *App) initModuleProvider() error {
	err := error(nil)
	app.moduleProvider, err = NewModuleProvider(app)
	if err != nil {
		app.logger.Errorf("%s", err.Error())
		return err
	}
	return nil
}

func (app *App) runHttpServer() error {
	if app.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			app.logger.Errorf("✖ Failed to load config: %s", err.Error())
			return fmt.Errorf("✖ Failed to load config: %v", err)
		}
		app.httpConfig = cfg
	}

	app.logger.Infof("🌐 Server is running on %s", app.httpConfig.Address())
	app.logger.Info("✅ Server started successfully")
	if err := app.app.Listen(app.httpConfig.Address()); err != nil {
		app.logger.Errorf("✖ Failed to start server: %s", err.Error())
		return fmt.Errorf("✖ Failed to start server: %v", err)
	}

	return nil
}

func (app *App) initRouter() error {
	app.app.Get("/swagger/*", swagger.HandlerDefault)
	api := app.app.Group("/api")

	app.moduleProvider.userModule.InitDelivery(api)
	return nil
}
