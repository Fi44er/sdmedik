package category

import (
	"strconv"

	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// Delete godoc
// @Summary Delete a category
// @Description Deletes a category
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} response.ResponseData "OK"
// @Router /category/{id} [delete]
func (i *Implementation) Delete(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := i.categoryService.Delete(ctx.Context(), id); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
