package middleware

import (
	"strings"

	"github.com/Fi44er/sdmedik/backend/config"
	"github.com/Fi44er/sdmedik/backend/module/user/dto"
	user_model "github.com/Fi44er/sdmedik/backend/module/user/model"
	"github.com/Fi44er/sdmedik/backend/shared/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func DeserializeUser(cache *redis.Client, db *gorm.DB, config *config.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var accessToken string
		authorization := ctx.Get("Authorization")
		userAgent := strings.ReplaceAll(ctx.Get("User-Agent"), " ", "")

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

		accessRedisKey := userAgent + ":" + tokenClaims.TokenUUID
		userID, err := cache.Get(ctx.Context(), accessRedisKey).Result()
		if err == redis.Nil {
			return ctx.Status(401).JSON(fiber.Map{
				"status":  "fail",
				"message": "Token is invalid or session has expired",
			})
		}

		var user user_model.User
		err = db.First(&user, "id = ?", userID).Error
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(403).JSON(fiber.Map{
				"status":  "fail",
				"message": "the user belonging to this token no logger exists",
			})
		}

		ctx.Locals("user", dto.FilterUserResponse(&user))
		ctx.Locals("access_token_uuid", tokenClaims.TokenUUID)

		return ctx.Next()
	}
}

func RoleRequired(roles ...string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user, ok := ctx.Locals("user").(*user_model.User)
		if !ok {
			return ctx.Status(403).JSON(fiber.Map{
				"status":  "fail",
				"message": "User  not found",
			})
		}

		for _, role := range roles {
			if user.Role == user_model.Role(role) {
				return ctx.Next() // Пользователь имеет нужную роль
			}
		}

		return ctx.Status(403).JSON(fiber.Map{
			"status":  "fail",
			"message": "You do not have permission to access this resource",
		})
	}
}
