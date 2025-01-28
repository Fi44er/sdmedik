package promotion

import (
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// Delete godoc
// @Summary Delete a promotion
// @Description Delete a promotion
// @Tags promotion
// @Produce json
// @Param id path string true "Promotion ID"
// @Success 200 {object} response.Response "OK"
// @Router /promotion/{id} [delete]
func (i *Implementation) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := i.promotionService.Delete(ctx.Context(), id); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}
	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
