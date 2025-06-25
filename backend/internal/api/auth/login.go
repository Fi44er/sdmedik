package auth

import (
	"strings"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	_ "github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// Login godoc
// @Summary Login user
// @Description Logs in a user and returns access and refresh tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.Login true "User  login credentials"
// @Success 200 {object} response.Response "OK"
// @Router /auth/login [post]
func (i *Implementation) Login(ctx *fiber.Ctx) error {
	user := new(dto.Login)
	var sessRes *session.Session
	sess := ctx.Locals("session")
	if sess != nil {
		sessRes = sess.(*session.Session)
	}
	userAgent := strings.ReplaceAll(ctx.Get("User-Agent"), " ", "")

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	accessToken, refreshToken, err := i.authService.Login(ctx.Context(), user, userAgent, sessRes)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	// expired := time.Now().Add(-time.Hour * 24)
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
