package order

import (
	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// Create godoc
// @Summary Create order
// @Description Creates a new order
// @Tags order
// @Accept json
// @Produce json
// @Param order body dto.CreateOrder true "Order details"
// @Success 200 {object} response.Response "OK"
// @Router /order [post]
func (i *Implementation) Create(ctx *fiber.Ctx) error {
	order := new(dto.CreateOrder)
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

	if err := ctx.BodyParser(order); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	url, err := i.orderService.Create(ctx.Context(), order, userRes.ID, sessRes)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": struct {
		URL string `json:"url"`
	}{URL: url}})
}
