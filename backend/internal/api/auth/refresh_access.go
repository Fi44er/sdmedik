package auth

import (
	_ "github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// RefreshAccessToken godoc
// @Summary Refresh access token
// @Description Refreshes the access token using the provided refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} response.Response "Successful refresh response"
// @Router /auth/refresh [post]
func (i *Implementation) RefreshAccessToken(ctx *fiber.Ctx) error {
	refreshToken := ctx.Cookies("refresh_token")
	accessToken, err := i.authService.RefreshAccessToken(ctx.Context(), refreshToken)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		MaxAge:   i.config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   i.config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: false,
	})

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
