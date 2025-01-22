package order

import (
	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// GetMyOrders godoc
// @Summary Get my orders
// @Description Get my orders
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseData "OK"
// @Router /order/my [get]
func (s *Implementation) GetMyOrders(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(response.UserResponse)
	orders, err := s.orderService.GetMyOrders(ctx.Context(), user.ID)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}
	return ctx.JSON(fiber.Map{"status": "success", "data": orders})
}
