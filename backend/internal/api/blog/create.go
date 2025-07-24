package blog

import (
	"github.com/Fi44er/sdmedik/backend/internal/model"
	_ "github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// Create godoc
// @Summary Create blog
// @Description Creating blog(ID указывать НЕ НУЖНО)
// @Tags blog
// @Accept json
// @Produce json
// @Param element body model.Blog true "Create blog"
// @Success 200 {object} response.Response "OK"
// @Router /blog [post]
func (i *Implementation) Create(ctx *fiber.Ctx) error {
	dto := new(model.Blog)

	if err := ctx.BodyParser(dto); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	if err := i.service.Create(ctx.Context(), dto); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}

// Update godoc
// @Summary Update blog
// @Description Updating blog(ID указывать НЕ НУЖНО)
// @Tags blog
// @Accept json
// @Produce json
// @Param element body model.Blog true "Update blog"
// @Success 200 {object} response.Response "OK"
// @Router /blog [put]
func (i *Implementation) Update(ctx *fiber.Ctx) error {
	dto := new(model.Blog)

	if err := ctx.BodyParser(dto); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	if err := i.service.Update(ctx.Context(), dto); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
