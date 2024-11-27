package user

import (
	"os"
	"strconv"
	"time"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

func (i *Implementation) Login(ctx *fiber.Ctx) error {
	user := new(dto.Login)

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	token, err := i.userService.Login(ctx.Context(), user)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	hour, _ := strconv.Atoi(os.Getenv("JWT_EXP"))

	ctx.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(time.Hour * time.Duration(hour)),
	})

	return ctx.Status(200).JSON("OK")
}
