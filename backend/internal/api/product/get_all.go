package product

import "github.com/gofiber/fiber/v2"

func (i *Implementation) GetAll(ctx *fiber.Ctx) error {
	offset := ctx.QueryInt("offset")
	limit := ctx.QueryInt("limit")

	products, err := i.productService.GetAll(ctx.Context(), offset, limit)
	if err != nil {
		return ctx.Status(400).JSON(err)
	}

	return ctx.Status(200).JSON(products)
}
