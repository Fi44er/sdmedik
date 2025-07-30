# Используем базовый образ Go на Alpine
FROM golang:1.23-alpine

# Устанавливаем GCC и зависимости для сборки
RUN apk update && apk add --no-cache build-base

# Копируем исходный код
WORKDIR /app
COPY . .

# Скачиваем зависимости и собираем приложение
RUN go mod download && \
    CGO_ENABLED=1 GOOS=linux go build -o main ./cmd

# Открываем порт (опционально)
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]
