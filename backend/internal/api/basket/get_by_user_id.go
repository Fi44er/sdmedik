package basket

import (
	"fmt"

	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// Get godoc
// @Summary Get basket
// @Description Get basket
// @Tags basket
// @Accept json
// @Produce json
// @Success 200 {object} response.Response "OK"
// @Router /basket [get]
func (i *Implementation) Get(ctx *fiber.Ctx) error {
	user := ctx.Locals("user")
	var userRes response.UserResponse
	var sessRes *session.Session
	if user != nil {
		userRes = user.(response.UserResponse)
	}
	sess := ctx.Locals("session")
	if sess != nil {
		sessRes = sess.(*session.Session)
	}

	fmt.Println(userRes.ID)
	basket, err := i.basketService.GetByUserID(ctx.Context(), userRes.ID, sessRes)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": basket})
}
