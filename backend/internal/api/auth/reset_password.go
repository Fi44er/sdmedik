package auth

import (
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// ResetPassword godoc
// @Summary Reset password
// @Description Resets the user's password
// @Tags auth
// @Accept json
// @Produce json
// @Param email path string true "User email"
// @Success 200 {object} response.Response "Successful password reset response"
// @Router /auth/reset-password/{email} [get]
func (i *Implementation) ResetPassword(ctx *fiber.Ctx) error {
	email := ctx.Params("email")
	if err := i.authService.ResetPassword(ctx.Context(), email); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}
	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}

// ChangePassword godoc
// @Summary Change password
// @Description Changes the user's password
// @Tags auth
// @Accept json
// @Produce json
// @Param token query string false "Reset password token"
// @Param pass query string false "New password"
// @Success 200 {object} response.Response "Successful password change response"
// @Router /auth/change-password [post]
func (i *Implementation) ChangePassword(ctx *fiber.Ctx) error {
	token := ctx.Query("token")
	pass := ctx.Query("pass")

	userID, err := i.authService.ValidateToken(ctx.Context(), token)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	if token == "" || pass == "" {
		return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
	}

	if err := i.authService.ChangePassword(ctx.Context(), token, pass, userID); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}
	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
