package page

import (
	"github.com/Fi44er/sdmedik/backend/internal/dto"
	_ "github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// AddElement godoc
// @Summary Add element to page
// @Description Adds element to page
// @Tags page
// @Accept json
// @Produce json
// @Param element body dto.AddElement true "Element to add"
// @Success 200 {object} response.Response "OK"
// @Router /page [post]
func (i *Implementation) AddElement(ctx *fiber.Ctx) error {
	dto := new(dto.AddElement)

	if err := ctx.BodyParser(dto); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	if err := i.service.AddElement(ctx.Context(), dto); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)

	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
