package search

import (
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// Search godoc
// @Summary Search products
// @Description Search products
// @Tags search
// @Accept json
// @Produce json
// @Param query query string false "Search query"
// @Success 200 {object} response.ResponseData "OK"
// @Router /search [get]
func (i *Implementation) Search(ctx *fiber.Ctx) error {
	query := ctx.Query("query")

	data, err := i.searchService.Search(ctx.Context(), query)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": data})
}
