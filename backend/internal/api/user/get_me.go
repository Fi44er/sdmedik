package user

import (
	"github.com/Fi44er/sdmedik/backend/internal/model"
	_ "github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/gofiber/fiber/v2"
)

// GetMy godoc
// @Summary Get my user
// @Description Get my user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} response.Response "OK"
// @Router /user/me [get]
func (i *Implementation) GetMy(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": user})
}
