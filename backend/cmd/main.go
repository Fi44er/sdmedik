package main

import (
	"github.com/Fi44er/sdmedik/backend/internal/app"
	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/pkg/database"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/Fi44er/sdmedik/backend/pkg/redis"
	"github.com/go-playground/validator/v10"
)

func main() {
	log := logger.GetLogger()
	validator := validator.New()

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

	_, err = redis.Connect(config.RedisUrl)
	if err != nil {
		log.Fatalf("✖ Connection error to redis: %v", err)
	}
	log.Info("🌏︎ Redis connected")

	a, err := app.NewApp(log, db.Db, validator, &config)
	if err != nil {
		log.Fatalf("✖ Failed to init app: %s", err.Error())
	}

	if err := a.Run(); err != nil {
		log.Fatalf("✖ Failed to run app: %s", err.Error())
	}
}
