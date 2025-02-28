package controller

import (
	"github.com/Fi44er/sdmedik/backend/module/product/converter"
	_ "github.com/Fi44er/sdmedik/backend/module/product/dto"
	_ "github.com/Fi44er/sdmedik/backend/shared/response"
	"github.com/gofiber/fiber/v2"
)

// GetByID godoc
// @Summary Get product by ID
// @Description Get product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} response.ResponseData{data=dto.ProductResponse} "OK"
// @Failure 500 {object} response.Response "Error"
// @Router /products/{id} [get]
func (c *ProductController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	product, err := c.service.GetByID(ctx.Context(), id)
	if err != nil {
		return err
	}

	response := converter.ToResponseFromDomain(product)

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": response})
}
