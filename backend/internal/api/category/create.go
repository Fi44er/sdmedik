package category

import (
	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// Create godoc
// @Summary Create a new category
// @Description Creates a new category
// @Tags category
// @Accept json
// @Produce json
// @Param category body dto.CreateCategory true "Category data"
// @Success 200 {object} response.ResponseData "OK"
// @Router /category [post]
func (i *Implementation) Create(ctx *fiber.Ctx) error {
	category := new(dto.CreateCategory)
	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	if err := i.categoryService.Create(ctx.Context(), category); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}