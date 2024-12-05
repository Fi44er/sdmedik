package product

import (
	_ "github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// Delete godoc
// @Summary Delete a product
// @Description Deletes a product
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} response.ResponseData "OK"
// @Router /product/{id} [delete]
func (i *Implementation) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := i.productService.Delete(ctx.Context(), id); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
