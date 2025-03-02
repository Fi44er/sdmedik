package controller

import (
	"github.com/Fi44er/sdmedik/backend/module/category/converter"
	_ "github.com/Fi44er/sdmedik/backend/module/category/dto"
	_ "github.com/Fi44er/sdmedik/backend/shared/response"
	"github.com/gofiber/fiber/v2"
)

// GetByID godoc
// @Summary Get category by ID
// @Description Get category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} response.ResponseData{data=dto.CategoryResponse} "OK"
// @Failure 500 {object} response.Response "Error"
// @Router /categories/{id} [get]
func (c *CategoryController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	category, err := c.service.GetByID(ctx.Context(), id)
	if err != nil {
		return err
	}

	response := converter.ToResponseFromDomain(category)

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": response})
}

// GetAll godoc
// @Summary Get all categories
// @Description Get all categories
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseData{data=[]dto.CategoryResponse} "OK"
// @Failure 500 {object} response.Response "Error"
// @Router /categories [get]
func (c *CategoryController) GetAll(ctx *fiber.Ctx) error {
	categories, err := c.service.GetAll(ctx.Context())
	if err != nil {
		return err
	}

	response := converter.ToResponseSliceFromDomainSlice(categories)

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": response})
}
