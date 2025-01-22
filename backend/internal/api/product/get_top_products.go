package product

import (
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// GetTopProducts godoc
// @Summary Get top products
// @Description Gets top products
// @Tags product
// @Accept json
// @Produce json
// @Param limit path integer true "Limit"
// @Success 200 {object} response.ResponseData "OK"
// @Router /product/top/{limit} [get]
func (i *Implementation) GetTopProducts(ctx *fiber.Ctx) error {
	limit, err := ctx.ParamsInt("limit")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status": "error", "message": "invalid limit"})
	}

	products, err := i.productService.GetTopProducts(ctx.Context(), limit)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}
	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": products})
}
