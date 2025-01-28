package main

import (
	_ "github.com/Fi44er/sdmedik/backend/docs"
	"github.com/Fi44er/sdmedik/backend/internal/app"
	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/pkg/database"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/Fi44er/sdmedik/backend/pkg/redis"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
	"github.com/go-playground/validator/v10"
)

//	@title				sdmedik API
//	@version		1.0
//	@description	Swagger docs from sdmedik backend
//	@host			127.0.0.1:8080
//	@BasePath		/api/v1/

func main() {
	log := logger.GetLogger()
	validator := validator.New()
	validator.RegisterValidation("characteristic_type", utils.CustomTypeValidator)

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("✖ Failed to load config: %s", err.Error())
	}
	log.Info("✔  Config loaded")

	db, err := database.ConnectDb(config.PostgresUrl)
	if err != nil {
		log.Fatalf("✖ Connection error to database: %v", err)
	}
	log.Info("🌏︎ Database connected")

	if err = database.Migrate(db.Db); err != nil {
		log.Fatalf("✖ Database migration failed: %v", err)
	}
	log.Info("✔  Database migrated")

	redis, err := redis.Connect(config.RedisUrl)
	if err != nil {
		log.Fatalf("✖ Connection error to redis: %v", err)
	}
	log.Info("🌏︎ Redis connected")

	a, err := app.NewApp(log, db.Db, validator, &config, redis.RedisClient)
	if err != nil {
		log.Fatalf("✖ Failed to init app: %s", err.Error())
	}

	if err := a.Run(); err != nil {
		log.Fatalf("✖ Failed to run app: %s", err.Error())
	}
}
