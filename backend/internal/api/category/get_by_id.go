package category

import (
	"strconv"

	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// GetByID godoc
// @Summary Get a category by ID
// @Description Gets a category by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} response.ResponseData "OK"
// @Router /category/{id} [get]
func (i *Implementation) GetByID(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	category, err := i.categoryService.GetByID(ctx.Context(), id)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": category})
}
