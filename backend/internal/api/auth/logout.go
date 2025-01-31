package auth

import (
	"strings"
	"time"

	_ "github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// Logout godoc
// @Summary Logout user
// @Description Logs out a user by clearing the access and refresh tokens
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} response.Response "Successful logout response"
// @Router /auth/logout [post]
func (i *Implementation) Logout(ctx *fiber.Ctx) error {
	refreshToken := ctx.Cookies("refresh_token")
	accessTokenUUID := ctx.Locals("access_token_uuid")
	userAgent := strings.ReplaceAll(ctx.Get("User-Agent"), " ", "")

	if err := i.authService.Logout(ctx.Context(), refreshToken, accessTokenUUID.(string), userAgent); err != nil {
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
