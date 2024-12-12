package user

import (
	_ "github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

// GetAll godoc
// @Summary Get all users
// @Description Get all users
// @Tags user
// @Accept json
// @Produce json
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Success 200 {object} response.Response "OK"
// @Router /user [get]
func (i *Implementation) GetAll(ctx *fiber.Ctx) error {
	offset := ctx.QueryInt("offset")
	limit := ctx.QueryInt("limit")

	users, err := i.userService.GetAll(ctx.Context(), offset, limit)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": users})
}
