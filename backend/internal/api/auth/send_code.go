package auth

import (
	"github.com/Fi44er/sdmedik/backend/internal/dto"
	_ "github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// SendCode godoc
// @Summary Send verification code
// @Description Sends a verification code to the provided email address
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.SendCode true "User email"
// @Success 200 {object} response.Response "Successful code sending response"
// @Router /send-code [post]
func (i *Implementation) SendCode(ctx *fiber.Ctx) error {
	data := new(dto.SendCode)

	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	if err := i.authService.SendCode(ctx.Context(), data.Email); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
