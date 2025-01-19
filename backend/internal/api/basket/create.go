package basket

import (
	"log"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// Create godoc
// @Summary Create a new basket
// @Description Create a new basket
// @Tags basket
// @Accept json
// @Produce json
// @Param basket body dto.CreateBasket true "Basket data"
// @Success 200 {object} response.Response "OK"
// @Router /basket/create [post]
func (i *Implementation) Create(ctx *fiber.Ctx) error {
	log.Println("create basket")
	basket := new(dto.CreateBasket)

	if err := ctx.BodyParser(basket); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	if err := i.basketService.Create(ctx.Context(), basket); err != nil {
		code, msg := errors.GetErroField(err)
		log.Println(err)
		return ctx.Status(code).JSON(msg)
	}
	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
