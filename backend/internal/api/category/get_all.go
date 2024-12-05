package category

import (
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// GetAll godoc
// @Summary Get all categories
// @Description Gets all categories
// @Tags category
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseListData "OK"
// @Router /category [get]
func (i *Implementation) GetAll(ctx *fiber.Ctx) error {
	categories, err := i.categoryService.GetAll(ctx.Context())
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": categories})
}
