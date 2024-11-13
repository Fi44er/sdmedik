package main

import (
	"github.com/Fi44er/sdmedik/backend/internal/app"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
)

func main() {
	log := logger.GetLogger()

	a, err := app.NewApp(log)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	if err := a.Run(); err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
