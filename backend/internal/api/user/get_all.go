package user

import (
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

func (i *Implementation) GetAll(ctx *fiber.Ctx) error {
	offset := ctx.QueryInt("offset")
	limit := ctx.QueryInt("limit")

	users, err := i.userService.GetAll(ctx.Context(), offset, limit)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}
	return ctx.Status(200).JSON(users)
}
