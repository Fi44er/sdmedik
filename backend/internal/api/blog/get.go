package blog

import (
	_ "github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// GetByID godoc
// @Summary Get blog by ID
// @Description Get single blog by its ID
// @Tags blog
// @Produce json
// @Param id path string true "Blog ID"
// @Success 200 {object} response.Response "OK"
// @Router /blog/{id} [get]
func (i *Implementation) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	blog, err := i.service.GetByID(ctx.Context(), id)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": blog})
}

// GetAll godoc
// @Summary Get all blogs
// @Description Get list of blogs with pagination
// @Tags blog
// @Produce json
// @Param offset query int false "Pagination offset" default(0)
// @Param limit query int false "Pagination limit" default(10)
// @Success 200 {object} response.Response "OK"
// @Router /blog [get]
func (i *Implementation) GetAll(ctx *fiber.Ctx) error {
	offset, _ := ctx.ParamsInt("offset")
	limit, _ := ctx.ParamsInt("limit")
	blogs, err := i.service.GetAll(ctx.Context(), offset, limit)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": blogs})
}
