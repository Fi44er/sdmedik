package product

import (
	"encoding/json"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// Update godoc
// @Summary Update a product
// @Description Updates a product with metadata (JSON) and multiple files
// @Description Example JSON:
// @Description ```
// @Description {
// @Description 	"name": "test update",
// @Description 	"price": 123.12,
// @Description 	"description": "test update",
// @Description 	"category_ids": [
// @Description 		2
// @Description 	],
// @Description 	"characteristic_values": [
// @Description 		{
// @Description 			"characteristic_id": 2,
// @Description 			"value": "string"
// @Description 		},
// @Description 		{
// @Description 			"characteristic_id": 3,
// @Description 			"value": "12"
// @Description 		}
// @Description 	],
// @Description 	"del_images": [
// @Description 		{
// @Description 			"id": "8832ed62-0a19-40e2-81b3-4c971884e962",
// @Description 			"name": "cb44cbb3-aaf9-47a4-867c-f81dbc0485ab.jpg"
// @Description 		}
// @Description 	]
// @Description }
// @Description ```
// @Tags product
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "Product ID"
// @Param json formData string true "Product metadata as JSON"
// @Param files formData []file false "Product images (multiple files)"
// @Success 200 {object} response.ResponseData "OK"
// @Router /product/{id} [put]
func (i *Implementation) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data := new(dto.UpdateProduct)

	jsonData := ctx.FormValue("json")
	if err := json.Unmarshal([]byte(jsonData), data); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status": "error", "message": "Failed to parse multipart form"})
	}

	files := form.File["files"]
	if len(files) > 5 {
		return ctx.Status(400).JSON(fiber.Map{"status": "error", "message": "Too many files"})
	}

	images := dto.Images{
		Files: files,
	}

	if err := i.productService.Update(ctx.Context(), data, &images, id); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
