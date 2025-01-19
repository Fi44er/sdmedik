package basket

import (
	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// Get godoc
// @Summary Get basket
// @Description Get basket
// @Tags basket
// @Accept json
// @Produce json
// @Success 200 {object} response.Response "OK"
// @Router /basket [get]
func (i *Implementation) Get(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(response.UserResponse)

	basket, err := i.basketService.GetByUserID(ctx.Context(), user.ID)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": basket})
}
