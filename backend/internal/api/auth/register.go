package auth

import (
	"github.com/Fi44er/sdmedik/backend/internal/dto"
	_ "github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// Register godoc
// @Summary User registration
// @Description Registers a new user with the provided data
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.Register true "User Register"
// @Success 200 {object} response.Response "Successful registration response"
// @Router /register [post]
func (i *Implementation) Register(ctx *fiber.Ctx) error {
	data := new(dto.Register)
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	if err := i.authService.Register(ctx.Context(), data); err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
