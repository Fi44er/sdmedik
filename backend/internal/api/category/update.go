package category

import (
	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// Update godoc
// @Summary Update a category
// @Tags category
// @Description Updates a category with metadata (JSON)
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body dto.UpdateCategory true "Category metadata as JSON"
// @Success 200 {object} response.ResponseData "OK"
// @Router /category/{id} [put]
func (i *Implementation) Update(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	dto := new(dto.UpdateCategory)
	if err := ctx.BodyParser(dto); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	if err := i.categoryService.Update(ctx.Context(), id, dto); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
