package basket

import (
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// CreateOrLoadGuestBasket godoc
// @Summary Create or load guest basket
// @Description Create or load guest basket
// @Tags basket
// @Accept json
// @Produce json
// @Param id path string true "Basket ID"
// @Success 200 {object} response.Response "OK"
// @Router /basket/guest/{id} [get]
func (i *Implementation) CreateOrLoadGuestBasket(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	basket, err := i.basketService.GetOrCreateGuestBasket(ctx.Context(), id)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": basket})
}
