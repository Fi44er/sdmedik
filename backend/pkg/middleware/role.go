package middleware

import (
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

// RoleRequired проверяет, имеет ли пользователь необходимую роль для доступа к эндпоинту.
func RoleRequired(roles ...string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user, ok := ctx.Locals("user").(*model.User)
		if !ok {
			return ctx.Status(403).JSON(fiber.Map{
				"status":  "fail",
				"message": "User  not found",
			})
		}

		for _, role := range roles {
			if user.Role.Name == role {
				return ctx.Next() // Пользователь имеет нужную роль
			}
		}

		return ctx.Status(403).JSON(fiber.Map{
			"status":  "fail",
			"message": "You do not have permission to access this resource",
		})
	}
}
