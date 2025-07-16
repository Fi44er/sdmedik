package page

import (
	_ "github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// GetByPath godoc
// @Summary Get page by path
// @Description Get page by path
// @Tags page
// @Accept json
// @Produce json
// @Success 200 {object} response.Response "OK"
// @Router /page/{path} [get]
func (i *Implementation) GetByPath(ctx *fiber.Ctx) error {
	path := ctx.Params("path")

	page, err := i.service.GetByPath(ctx.Context(), path)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": page})
}
