package webscraper

import "github.com/gofiber/fiber/v2"

// Scraper godoc
// @Summary Scraper
// @Description Scraper
// @Tags webscraper
// @Accept json
// @Produce json
// @Success 200 {object} response.Response "OK"
// @Router /webscraper [get]
func (i *Implementation) Scraper(ctx *fiber.Ctx) error {
	if err := i.webscraperService.Scraper(); err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}
