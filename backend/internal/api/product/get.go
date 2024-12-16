package product

import (
	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (i *Implementation) Get(ctx *fiber.Ctx) error {
	params := ctx.Queries()
	var criteria dto.ProductSearchCriteria

	utils.BindQueryToStruct(params, &criteria)

	product, err := i.productService.Get(ctx.Context(), criteria)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	if len(product) == 1 {
		return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": product[0]})
	}
	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": product})
}
