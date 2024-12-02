package auth

import (
	"github.com/Fi44er/sdmedik/backend/internal/dto"
	_ "github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// VerifyCode godoc
// @Summary Verify the provided code
// @Description Verifies the code sent to the user's email
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.VerifyCode true "User verification code"
// @Success 200 {object} response.Response "Successful verification response"
// @Router /verify-code [post]
func (i *Implementation) VerifyCode(ctx *fiber.Ctx) error {
	data := new(dto.VerifyCode)
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	if err := i.authService.VerifyCode(ctx.Context(), data); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
