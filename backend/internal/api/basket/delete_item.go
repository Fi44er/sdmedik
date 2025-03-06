package basket

import (
	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
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
	user := ctx.Locals("user")
	var userRes response.UserResponse
	var sessRes *session.Session
	if user != nil {
		userRes = user.(response.UserResponse)
	}
	sess := ctx.Locals("session")
	if sess != nil {
		sessRes = sess.(*session.Session)
	}

	if err := i.basketService.DeleteItem(ctx.Context(), itemID, userRes.ID, sessRes); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}
	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
