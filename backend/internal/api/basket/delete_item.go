package basket

import (
	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// DeleteItem godoc
// @Summary Delete item from basket
// @Description Delete item from basket
// @Tags basket
// @Accept json
// @Produce json
// @Param id path string true "Item ID"
// @Success 200 {object} response.Response "OK"
// @Router /basket/{id} [delete]
func (i *Implementation) DeleteItem(ctx *fiber.Ctx) error {
	itemID := ctx.Params("id")
	user := ctx.Locals("user").(response.UserResponse)

	if err := i.basketService.DeleteItem(ctx.Context(), itemID, user.ID); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}
	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
