package controller

import (
	"encoding/json"

	"github.com/Fi44er/sdmedik/backend/module/product/converter"
	"github.com/Fi44er/sdmedik/backend/module/product/dto"
	_ "github.com/Fi44er/sdmedik/backend/shared/response"
	"github.com/Fi44er/sdmedik/backend/shared/utils"
	"github.com/gofiber/fiber/v2"
)

// Create Product
// @Summary      Create product
// @Description  Create product
// @Tags         products
// @Accept       multipart/form-data
// @Produce      json
// @Param        product  formData  string  true  "A JSON object with product information (as a string)"
// @Param        files    formData  file    false "Image files for the product (multiple possible)"
// @Success      200      {object}  response.Response "OK"
// @Failure      500      {object}  response.Response "Error"
// @Router       /products [post]
func (c *ProductController) Create(ctx *fiber.Ctx) error {
	productDTO := new(dto.CreateProductDTO)
	productJson := ctx.FormValue("product")
	if err := json.Unmarshal([]byte(productJson), productDTO); err != nil {
		c.logger.Errorf("error parsing JSON: %s", err)
		return ctx.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid JSON format"})
	}

	domain, err := utils.ParseAndValidate(ctx, productDTO, c.validator, converter.ToDomainFromDTO, c.logger)
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
		c.logger.Errorf("error while creating product: %s", err)
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to create product"})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Product created successfully",
	})
}
