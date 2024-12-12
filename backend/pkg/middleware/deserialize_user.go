package middleware

import (
	"strings"

	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func DeserializeUser(cache *redis.Client, db *gorm.DB, config *config.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var accessToken string
		authorization := ctx.Get("Authorization")

		if strings.HasPrefix(authorization, "Bearer ") {
			accessToken = strings.TrimPrefix(authorization, "Bearer ")
		} else if ctx.Cookies("access_token") != "" {
			accessToken = ctx.Cookies("access_token")
		}

		if accessToken == "" {
			return ctx.Status(401).JSON(fiber.Map{
				"status":  "fail",
				"message": "You are not logged in",
			})
		}

		tokenClaims, err := utils.ValidateToken(accessToken, config.AccessTokenPublicKey)
		if err != nil {
			return ctx.Status(401).JSON(fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}

		userID, err := cache.Get(ctx.Context(), tokenClaims.TokenUUID).Result()
		if err == redis.Nil {
			return ctx.Status(401).JSON(fiber.Map{
				"status":  "fail",
				"message": "Token is invalid or session has expired",
			})
		}

		var user model.User
		err = db.Preload("Role").First(&user, "id = ?", userID).Error
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(403).JSON(fiber.Map{
				"status":  "fail",
				"message": "the user belonging to this token no logger exists",
			})
		}

		ctx.Locals("user", response.FilterUserResponse(&user))
		ctx.Locals("access_token_uuid", tokenClaims.TokenUUID)

		return ctx.Next()
	}
}
