package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var e *logrus.Entry

// Logger - это структура логгера
type Logger struct {
	*logrus.Entry
}

// GetLogger возвращает экземпляр логгера
func GetLogger() Logger {
	return Logger{e}
}

// NewLogger создает новый экземпляр логгера с настройками
func init() {
	log := logrus.New()

	// Устанавливаем формат вывода
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05", // Изменяем формат времени
	})

	// Устанавливаем уровень логирования
	log.SetLevel(logrus.DebugLevel)

	// Устанавливаем вывод в консоль
	log.SetOutput(os.Stdout)

	e = logrus.NewEntry(log)
}

// Error добавляет информацию о месте вызова и логирует сообщение об ошибке
func (l Logger) Error(msg string) {
	l.logWithCaller(logrus.ErrorLevel, msg)
}

// Info добавляет информацию о месте вызова и логирует информационное сообщение
func (l Logger) Info(msg string) {
	l.logWithCaller(logrus.InfoLevel, msg)
}

// Warn добавляет информацию о месте вызова и логирует предупреждение
func (l Logger) Warn(msg string) {
	l.logWithCaller(logrus.WarnLevel, msg)
}

// logWithCaller добавляет информацию о вызове
func (l Logger) logWithCaller(level logrus.Level, msg string) {
	_, file, line, ok := runtime.Caller(2) // Используем 2, чтобы получить информацию о вызове логгера
	if ok {
		fileLine := fmt.Sprintf("%s:%d", path.Base(file), line)

		// Добавляем ANSI-коды для цвета (например, зеленый)

		coloredFileLine := fmt.Sprintf("\033[32m%s\033[0m", fileLine) // Зеленый цвет

		l.Entry.Log(level, fmt.Sprintf("%s %s", coloredFileLine, msg))
	} else {
		l.Entry.Log(level, msg)
	}
}
