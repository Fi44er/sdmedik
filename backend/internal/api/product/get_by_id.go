package product

import (
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// GetByID godoc
// @Summary Get a product by ID
// @Description Gets a product by ID
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} response.ResponseData "OK"
// @Router /product/{id} [get]
func (i *Implementation) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	product, err := i.productService.GetByID(ctx.Context(), id)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": product})
}
