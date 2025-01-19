package basket

import (
	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// AddItem godoc
// @Summary Add item to basket
// @Description Add item to basket
// @Tags basket
// @Accept json
// @Produce json
// @Param basket body dto.AddBasketItem true "Basket item data"
// @Success 200 {object} response.Response "OK"
// @Router /basket [post]
func (i *Implementation) AddItem(ctx *fiber.Ctx) error {
	basketItem := new(dto.AddBasketItem)
	user := ctx.Locals("user").(response.UserResponse)

	if err := ctx.BodyParser(basketItem); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	if err := i.basketService.AddItem(ctx.Context(), basketItem, user.ID); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
