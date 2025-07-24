package blog

import (
	"io"

	_ "github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// Upload handles file upload
// @Summary Upload a file
// @Description Uploads a file to the server
// @Tags blog
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 200 {object} response.Response "OK"
// @Router /blog/upload [post]
func (i *Implementation) Upload(ctx *fiber.Ctx) error {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "File is required",
		})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to read file",
		})
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to read file content",
		})
	}

	link, err := i.service.Upload(ctx.Context(), fileHeader.Filename, fileData)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)

	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": link})
}
