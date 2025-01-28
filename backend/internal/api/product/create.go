package product

import (
	"encoding/json"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	_ "github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// Product godoc
// @Summary Create a new product
// @Description Creates a new product with metadata (JSON) and multiple files.
// @Description Example JSON:
// @Description ```
// @Description {
// @Description 	"article": "123-123-124",
// @Description 	"category_ids": [
// @Description 		2
// @Description 	],
// @Description 	"characteristic_values": [
// @Description 		{
// @Description 			"characteristic_id": 3,
// @Description 			"value": "12"
// @Description 		}
// @Description 	],
// @Description 	"description": "description #1",
// @Description 	"name": "product #1",
// @Description 	"price": 123.12
// @Description }
// @Description ```
// @Tags product
// @Accept multipart/form-data
// @Produce json
// @Param json formData string true "Product metadata as JSON"
// @Param files formData []file true "Product images (multiple files)"
// @Success 200 {object} response.ResponseData "OK"
// @Router /product [post]
func (i *Implementation) Create(ctx *fiber.Ctx) error {
	product := new(dto.CreateProduct)

	jsonData := ctx.FormValue("json")
	if err := json.Unmarshal([]byte(jsonData), product); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status": "error", "message": "Failed to parse multipart form"})
	}

	// Извлекаем массив файлов из поля "files"
	files := form.File["files"]
	if len(files) > 5 {
		return ctx.Status(400).JSON(fiber.Map{"status": "error", "message": "Too many files"})
	}

	images := dto.Images{
		Files: files,
	}

	if err := i.productService.Create(ctx.Context(), product, &images); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
