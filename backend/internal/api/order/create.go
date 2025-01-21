package order

import (
	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// Create godoc
// @Summary Create order
// @Description Creates a new order
// @Tags order
// @Accept json
// @Produce json
// @Param order body dto.CreateOrder true "Order details"
// @Success 200 {object} response.Response "OK"
// @Router /order [post]
func (i *Implementation) Create(ctx *fiber.Ctx) error {
	order := new(dto.CreateOrder)
	user := ctx.Locals("user").(response.UserResponse)

	if err := ctx.BodyParser(order); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	url, err := i.orderService.Create(ctx.Context(), order, user.ID)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": struct {
		ID string `json:"id"`
	}{ID: url}})
}
