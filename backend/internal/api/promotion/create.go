package promotion

import (
	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// Promotion godoc
// @Summary Create a new promotion
// @Description Creates a new promotion
// @Tags promotion
// @Accept json
// @Produce json
// @Param promotion body dto.CreatePromotion true "Promotion data"
// @Success 200 {object} response.Response "OK"
// @Router /promotion [post]
func (i *Implementation) Create(ctx *fiber.Ctx) error {
	data := new(dto.CreatePromotion)

	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	if err := i.promotionService.Create(ctx.Context(), data); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
