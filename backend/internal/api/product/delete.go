package product

import (
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

func (i *Implementation) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := i.productService.Delete(ctx.Context(), id); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}
	return ctx.Status(200).JSON("OK")
}
