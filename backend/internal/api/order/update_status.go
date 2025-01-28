package order

import (
	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// UpdateStatus godoc
// @Summary Update order status
// @Description Update order status
// @Tags order
// @Accept json
// @Produce json
// @Param order body dto.ChangeOrderStatus true "Order status data"
// @Success 200 {object} response.Response "OK"
// @Router /order/status [put]
func (i *Implementation) UpdateStatus(ctx *fiber.Ctx) error {
	data := new(dto.ChangeOrderStatus)

	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	if err := i.orderService.ChangeStatus(ctx.Context(), data); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
