package product

import (
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (i *Implementation) Create(ctx *fiber.Ctx) error {
	product := new(model.Product)
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	if err := i.productService.Create(ctx.Context(), product); err != nil {
		return ctx.Status(400).JSON(err)
	}
	return nil
}
