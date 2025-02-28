package controller

import (
	"encoding/json"

	"github.com/Fi44er/sdmedik/backend/module/category/converter"
	"github.com/Fi44er/sdmedik/backend/module/category/dto"
	"github.com/Fi44er/sdmedik/backend/shared/utils"
	"github.com/gofiber/fiber/v2"
)

// CreateCategory godoc
// @Summary      Create a new category
// @Description  Creates a new category with characteristics and images
// @Tags         categories
// @Accept       multipart/form-data
// @Produce      json
// @Param        category  formData  string  true  "Category data in JSON format"
// @Param        files     formData  file    false "Category images (can be multiple)"
// @Router       /categories [post]
func (c *CategoryController) Create(ctx *fiber.Ctx) error {
	categoryDTO := new(dto.CreateCategoryDTO)
	categoryJSON := ctx.FormValue("category")
	if err := json.Unmarshal([]byte(categoryJSON), categoryDTO); err != nil {
		c.logger.Errorf("error parsing JSON: %s", err)
		return ctx.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid JSON format"})
	}

	domain, err := utils.ParseAndValidate(ctx, categoryDTO, c.validator, converter.ToDomainFromDTO, c.logger)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		c.logger.Errorf("error while parsing multipart form: %s", err)
		return ctx.Status(400).JSON(fiber.Map{"status": "error", "message": "Failed to parse multipart form"})
	}
	files := form.File["files"]

	if err := c.service.Create(ctx.Context(), domain, files); err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Category created successfully"})
}
