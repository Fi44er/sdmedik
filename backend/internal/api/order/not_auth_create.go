package order

import (
	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// NotAuthCreate godoc
// @Summary Create order
// @Description Creates a new order
// @Tags order
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param order body dto.CreateOrder true "Order details"
// @Success 200 {object} response.Response "OK"
// @Router /order/{id} [post]
func (i *Implementation) NotAuthCreate(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data := new(dto.CreateOrder)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}
	url, err := i.orderService.NotAuthCreate(ctx.Context(), data, id)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": struct {
		URL string `json:"url"`
	}{
		URL: url,
	}})
}
