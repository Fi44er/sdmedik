package product

import (
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

func (i *Implementation) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	product := new(model.Product)
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	product.ID = id
	if err := i.productService.Update(ctx.Context(), product); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
