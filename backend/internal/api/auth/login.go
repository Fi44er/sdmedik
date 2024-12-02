package auth

import (
	"github.com/Fi44er/sdmedik/backend/internal/dto"
	_ "github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// Login godoc
// @Summary Login user
// @Description Logs in a user and returns access and refresh tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.Login true "User  login credentials"
// @Success 200 {object} response.Response "OK"
// @Router /login [post]
func (i *Implementation) Login(ctx *fiber.Ctx) error {
	user := new(dto.Login)

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	accessToken, refreshToken, err := i.authService.Login(ctx.Context(), user)
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
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		MaxAge:   i.config.RefreshTokenMaxAge * 60,
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
