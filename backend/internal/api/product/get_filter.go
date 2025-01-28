package product

import (
	_ "github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// GetFilter godoc
// @Summary Get a product
// @Description Get a product filter
// @Tags product
// @Accept json
// @Produce json
// @Param category_id path string true "Category ID"
// @Success 200 {object} response.ResponseData "OK"
// @Router /product/filter/{category_id} [get]
func (h *Implementation) GetFilter(ctx *fiber.Ctx) error {
	categoryID, _ := ctx.ParamsInt("category_id")
	filter, err := h.productService.GetFilter(ctx.Context(), categoryID)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}
	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": filter})
}
