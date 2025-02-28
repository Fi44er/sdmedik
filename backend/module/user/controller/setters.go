package controller

import (
	"github.com/Fi44er/sdmedik/backend/module/user/converter"
	"github.com/Fi44er/sdmedik/backend/module/user/dto"
	_ "github.com/Fi44er/sdmedik/backend/shared/response"
	"github.com/Fi44er/sdmedik/backend/shared/utils"
	"github.com/gofiber/fiber/v2"
)

// Create godoc
// @Summary Create user
// @Description Create user
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.UserDTO true "User"
// @Success      200      {object}  response.Response "OK"
// @Failure      500      {object}  response.Response "Error"
// @Router /users [post]
func (c *UserController) Create(ctx *fiber.Ctx) error {
	dto := new(dto.UserDTO)

	domain, err := utils.ParseAndValidate(ctx, dto, c.validator, converter.ToDomainFromDto, c.logger)
	if err != nil {
		return err
	}

	if err := c.service.Create(ctx.Context(), domain); err != nil {
		c.logger.Errorf("error while create user: %s", err)
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "user created successfully",
	})
}

// Update godoc
// @Summary Update user
// @Description Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body dto.UserDTO true "User"
// @Success      200      {object}  response.Response "OK"
// @Failure      500      {object}  response.Response "Error"
// @Router /users/{id} [put]
func (c *UserController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	dto := new(dto.UserDTO)

	domain, err := utils.ParseAndValidate(ctx, dto, c.validator, converter.ToDomainFromDto, c.logger)
	if err != nil {
		return err
	}
	domain.ID = id

	if err := c.service.Update(ctx.Context(), domain); err != nil {
		c.logger.Errorf("error while update user: %s", err)
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "user updated successfully",
	})
}

// Delete godoc
// @Summary Delete user
// @Description Delete user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success      200      {object}  response.Response "OK"
// @Failure      500      {object}  response.Response "Error"
// @Router /users/{id} [delete]
func (c *UserController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := c.service.Delete(ctx.Context(), id); err != nil {
		c.logger.Errorf("error while delete user: %s", err)
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "user deleted successfully",
	})
}
