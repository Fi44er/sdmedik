package promotion

import (
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// GetAll godoc
// @Summary Get all promotions
// @Description Get all promotions
// @Tags promotion
// @Produce json
// @Success 200 {object} response.Response "OK"
// @Router /promotion [get]
func (i *Implementation) GetAll(ctx *fiber.Ctx) error {
	promotions, err := i.promotionService.GetAll(ctx.Context())
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}
	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": promotions})
}
