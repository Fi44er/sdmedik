package product

import (
	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

func (i *Implementation) Create(ctx *fiber.Ctx) error {
	product := new(dto.CreateProduct)
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	if err := i.productService.Create(ctx.Context(), product); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}