package chat

import (
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// GetAll godoc
// @Summary Get all chats
// @Description Get all chats
// @Tags chat
// @Accept json
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} response.ResponseData "OK"
// @Router /chat [get]
func (i *Implementation) GetAll(ctx *fiber.Ctx) error {
	limit := ctx.QueryInt("limit")
	offset := ctx.QueryInt("offset")

	chat, err := i.service.GetAll(ctx.Context(), limit, offset)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}
	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": chat})
}
