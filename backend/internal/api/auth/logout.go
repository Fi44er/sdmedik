package auth

import (
	"time"

	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

func (i *Implementation) Logout(ctx *fiber.Ctx) error {
	refreshToken := ctx.Cookies("refresh_token")
	accessTokenUUID := ctx.Locals("access_token_uuid")

	if err := i.authService.Logout(ctx.Context(), refreshToken, accessTokenUUID.(string)); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	expired := time.Now().Add(-time.Hour * 24)
	ctx.Cookie(&fiber.Cookie{
		Name:    "access_token",
		Value:   "",
		Expires: expired,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:    "refresh_token",
		Value:   "",
		Expires: expired,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:    "logged_in",
		Value:   "",
		Expires: expired,
	})

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
