package chat

import (
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// DeleteChat godoc
// @Summary Delete chat
// @Description Delete chat
// @Tags chat
// @Accept json
// @Produce json
// @Param id path string true "Chat ID"
// @Success 200 {object} response.Response "OK"
// @Router /chat/{id} [delete]
func (i *Implementation) DeleteChat(ctx *fiber.Ctx) error {
	chatID := ctx.Params("id")
	err := i.service.DeleteChat(ctx.Context(), chatID)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}
	return ctx.Status(200).JSON(fiber.Map{"status": "success"})
}
