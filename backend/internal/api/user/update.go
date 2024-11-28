package user

import (
	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

func (i *Implementation) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data := new(dto.UpdateUser)
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	if err := i.userService.Update(ctx.Context(), data, id); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
