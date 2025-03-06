package middleware

import (
	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func AllowGuest(cache *redis.Client, db *gorm.DB, config *config.Config, store *session.Store) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		_ = DeserializeUser(cache, db, config)(ctx)
		sess, err := store.Get(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Session error",
			})
		}
		ctx.Locals("session", sess)

		return ctx.Next()
	}
}
