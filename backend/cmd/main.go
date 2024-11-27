package main

import (
	"github.com/Fi44er/sdmedik/backend/internal/app"
	"github.com/Fi44er/sdmedik/backend/pkg/database"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/Fi44er/sdmedik/backend/pkg/redis"
	"github.com/go-playground/validator/v10"
)

func main() {
	log := logger.GetLogger()
	validator := validator.New()

	db, err := database.ConnectDb()
	if err != nil {
		log.Fatalf("Connection error to database: %v", err)
	}

	if err = database.Migrate(db.Db); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	_, err = redis.Connect()
	if err != nil {
		log.Fatalf("Connection error to redis: %v", err)
	}

	a, err := app.NewApp(log, db.Db, validator)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	if err := a.Run(); err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
