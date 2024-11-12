package main

import (
	"log"

	"github.com/Fi44er/sdmedik/backend/internal/app"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
)

func main() {
	log1 := logger.GetLogger()

	log1.Info("Информационное сообщение")

	log1.Warn("Предупреждение")

	log1.Error("Ошибка")

	// Пример использования с полями

	a, err := app.NewApp()
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	if err := a.Run(); err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
