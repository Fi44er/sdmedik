package controller

import (
	"strings"
	"time"

	"github.com/Fi44er/sdmedik/backend/module/auth/dto"
	_ "github.com/Fi44er/sdmedik/backend/shared/response"
	"github.com/gofiber/fiber/v2"
)

// Register godoc
// @Summary User registration
// @Description Registers a new user with the provided data
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.RegisterDTO true "User Register"
// @Success      200      {object}  response.Response "OK"
// @Failure      500      {object}  response.Response "Error"
// @Router /auth/register [post]
func (c *AuthController) Register(ctx *fiber.Ctx) error {
	data := new(dto.RegisterDTO)
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	if err := c.authServ.Register(ctx.Context(), data); err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}

// Logout godoc
// @Summary Logout user
// @Description Logs out a user by clearing the access and refresh tokens
// @Tags auth
// @Accept json
// @Produce json
// @Success      200      {object}  response.Response "OK"
// @Failure      500      {object}  response.Response "Error"
// @Router /auth/logout [post]
func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	refreshToken := ctx.Cookies("refresh_token")
	accessTokenUUID := ctx.Locals("access_token_uuid")

	dto := dto.LogoutDTO{
		RefreshToken:    refreshToken,
		AccessTokenUUID: accessTokenUUID.(string),
	}

	if err := c.authServ.Logout(ctx.Context(), &dto); err != nil {
		return err
	}

	expired := time.Now().Add(-time.Hour * 24)
	ctx.Cookie(&fiber.Cookie{
		Name:    "access_token",
		Value:   "",
		Expires: expired,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:    "refresh_token",
		Value:   "",
		Expires: expired,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:    "logged_in",
		Value:   "",
		Expires: expired,
	})

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}

// RefreshAccessToken godoc
// @Summary Refresh access token
// @Description Refreshes the access token using the provided refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Success      200      {object}  response.Response "OK"
// @Failure      500      {object}  response.Response "Error"
// @Router /auth/refresh [post]
func (c *AuthController) RefreshAccessToken(ctx *fiber.Ctx) error {
	refreshToken := ctx.Cookies("refresh_token")
	userAgent := strings.ReplaceAll(ctx.Get("User-Agent"), " ", "")

	dto := dto.RefreshTokenDTO{
		RefreshToken: refreshToken,
		UserAgent:    userAgent,
	}
	accessToken, err := c.authServ.RefreshAccessToken(ctx.Context(), &dto)
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		MaxAge:   c.config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   c.config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: false,
	})

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}

// Login godoc
// @Summary Login user
// @Description Logs in a user and returns access and refresh tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.LoginDTO true "User  login credentials"
// @Success      200      {object}  response.Response "OK"
// @Failure      500      {object}  response.Response "Error"
// @Router /auth/login [post]
func (c *AuthController) Login(ctx *fiber.Ctx) error {
	dto := new(dto.LoginDTO)
	userAgent := strings.ReplaceAll(ctx.Get("User-Agent"), " ", "")

	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	dto.UserAgent = userAgent

	loginRes, err := c.authServ.Login(ctx.Context(), dto)
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    loginRes.AccessToken,
		Path:     "/",
		MaxAge:   c.config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    loginRes.RefreshToken,
		Path:     "/",
		MaxAge:   c.config.RefreshTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   c.config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: false,
	})

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}

// SendCode godoc
// @Summary Send verification code
// @Description Sends a verification code to the provided email address
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.SendCodeDTO true "User email"
// @Success      200      {object}  response.Response "OK"
// @Failure      500      {object}  response.Response "Error"
// @Router /auth/send-code [post]
func (c *AuthController) SendCode(ctx *fiber.Ctx) error {
	data := new(dto.SendCodeDTO)

	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	if err := c.authServ.SendCode(ctx.Context(), data.Email); err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}

// VerifyCode godoc
// @Summary Verify the provided code
// @Description Verifies the code sent to the user's email
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.VerifyCodeDTO true "User verification code"
// @Success      200      {object}  response.Response "OK"
// @Failure      500      {object}  response.Response "Error"
// @Router /auth/verify-code [post]
func (c *AuthController) VerifyCode(ctx *fiber.Ctx) error {
	data := new(dto.VerifyCodeDTO)
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(400).JSON("Failed to parse body")
	}

	if err := c.authServ.VerifyCode(ctx.Context(), data); err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
