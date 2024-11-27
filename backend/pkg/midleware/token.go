package midleware

import (
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func JWTProtected(ctx *fiber.Ctx) error {
	token := ctx.Cookies("token")

	if token == "" {
		return errors.New(401, "Unauthorized")
	}

	bool, err := utils.VerifyToken(token)
	if err != nil || !bool {
		return errors.New(401, "Unauthorized")
	}

	return ctx.Next()
}
