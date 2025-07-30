package blog

import (
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// Delete godoc
// @Summary Delete blog by ID
// @Description Delete blog by its ID
// @Tags blog
// @Produce json
// @Param id path string true "Blog ID"
// @Success 200 {object} response.Response "OK"
// @Router /blog/{id} [delete]
func (i *Implementation) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := i.service.Delete(ctx.Context(), id)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}
	return ctx.Status(200).JSON(fiber.Map{"status": "success"})
}
