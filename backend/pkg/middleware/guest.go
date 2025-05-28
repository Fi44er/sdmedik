package middleware

import (
	"strings"

	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func AllowGuest(cache *redis.Client, db *gorm.DB, config *config.Config, store *session.Store) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var accessToken string
		var err error
		authorizeStatus := true

		authorization := ctx.Get("Authorization")
		userAgent := strings.ReplaceAll(ctx.Get("User-Agent"), " ", "")

		if strings.HasPrefix(authorization, "Bearer ") {
			accessToken = strings.TrimPrefix(authorization, "Bearer ")
		} else if ctx.Cookies("access_token") != "" {
			accessToken = ctx.Cookies("access_token")
		}

		if accessToken == "" {
			authorizeStatus = false
		}

		var tokenClaims *utils.TokenDetails
		if authorizeStatus {
			tokenClaims, err = utils.ValidateToken(accessToken, config.AccessTokenPublicKey)
			if err != nil {
				authorizeStatus = false
			}
		}

		var userID string
		if authorizeStatus {
			accessRedisKey := userAgent + ":" + tokenClaims.TokenUUID
			userID, err = cache.Get(ctx.Context(), accessRedisKey).Result()
			if err == redis.Nil {
				authorizeStatus = false
			}
		}

		var user model.User
		if authorizeStatus {
			err = db.Preload("Role").First(&user, "id = ?", userID).Error
			if err == gorm.ErrRecordNotFound {
				authorizeStatus = false
			}
		}

		var sess *session.Session
		if authorizeStatus {
			ctx.Locals("user", response.FilterUserResponse(&user))
			ctx.Locals("access_token_uuid", tokenClaims.TokenUUID)
		} else {
			sess, err = store.Get(ctx)
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Session error",
				})
			}

			ctx.Locals("session", sess)
			ctx.Locals("session_id", sess.ID())
		}

		_ = ctx.Next()
		return sess.Save()
	}
}
