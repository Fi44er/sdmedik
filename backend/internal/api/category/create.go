package category

import (
	"encoding/json"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// Create godoc
// @Summary Create a new category
// @Description Creates a new category with metadata (JSON) and a file (image)
// @Description Example JSON:
// @Description ```
// @Description {
// @Description 	"name": "category #1",
// @Description 	"characteristics": [
// @Description 		{
// @Description 			"data_type": "string",
// @Description 			"name": "characteristic #1"
// @Description 		},
// @Description 		{
// @Description 			"data_type": "int",
// @Description 			"name": "characteristic #2"
// @Description 		}
// @Description 	]
// @Description }
// @Tags category
// @Accept multipart/form-data
// @Produce json
// @Param json formData string true "Category metadata as JSON"
// @Param file formData file true "Category image file"
// @Success 200 {object} response.ResponseData "OK"
// @Router /category [post]
func (i *Implementation) Create(ctx *fiber.Ctx) error {
	category := new(dto.CreateCategory)
	jsonData := ctx.FormValue("json")
	if err := json.Unmarshal([]byte(jsonData), category); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	image := dto.Image{
		File: file,
	}

	if err := i.categoryService.Create(ctx.Context(), category, &image); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
