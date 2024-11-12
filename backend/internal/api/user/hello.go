package user

import (
	"github.com/gofiber/fiber/v2"
)

func (i *Implementation) Hello(ctx *fiber.Ctx) error {
	string := i.userService.Hello(ctx.Context())
	return ctx.Status(200).JSON(string)
}
