package auth

import (
	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

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
