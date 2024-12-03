package user

import (
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (i *Implementation) GetMy(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": user})
}
