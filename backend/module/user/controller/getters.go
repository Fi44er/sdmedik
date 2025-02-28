package controller

import (
	"github.com/Fi44er/sdmedik/backend/module/user/converter"
	"github.com/Fi44er/sdmedik/backend/module/user/dto"
	_ "github.com/Fi44er/sdmedik/backend/shared/response"
	"github.com/gofiber/fiber/v2"
)

// GetMy godoc
// @Summary Get my user
// @Description Get my user
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseData{data=dto.UserResponse} "OK"
// @Failure 500 {object} response.Response "Error"
// @Router /users/me [get]
func (c *UserController) GetMy(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(dto.UserResponse)
	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": user})
}

// GetByID godoc
// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.ResponseData{data=dto.UserResponse} "OK"
// @Failure 500 {object} response.Response "Error"
// @Router /users/{id} [get]
func (c *UserController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user, err := c.service.GetByID(ctx.Context(), id)
	if err != nil {
		return err
	}

	response := converter.ToResponseFromDomain(user)

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": response})
}

// GetAll godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Success 200 {object} response.ResponseListData{data=[]dto.UserResponse} "OK"
// @Failure 500 {object} response.Response "Error"
// @Router /users [get]
func (c *UserController) GetAll(ctx *fiber.Ctx) error {
	offset := ctx.QueryInt("offset")
	limit := ctx.QueryInt("limit")

	users, err := c.service.GetAll(ctx.Context(), offset, limit)
	if err != nil {
		return err
	}

	response := converter.ToResponseSliceFromDomain(users)

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": response})
}
