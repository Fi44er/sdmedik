package webscraper

import "github.com/gofiber/fiber/v2"

// Scraper godoc
// @Summary Scraper
// @Description Scraper
// @Tags webscraper
// @Accept json
// @Produce json
// @Success 200 {object} response.Response "OK"
// @Router /webscraper/start [post]
func (i *Implementation) Scraper(ctx *fiber.Ctx) error {
	if err := i.webscraperService.Scraper(); err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "OK"})
}

// CancelScraper godoc
// @Summary Отмена работы парсера
// @Description Отменяет выполнение текущего парсинга
// @Tags webscraper
// @Accept json
// @Produce json
// @Success 200 {object} response.Response "Парсинг отменён"
// @Router /webscraper/cancel [post]
func (i *Implementation) CancelScraper(ctx *fiber.Ctx) error {
	err := i.webscraperService.CancelScraper()
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status": "error", "message": "Нет активного парсинга"})
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "Парсинг отменён"})
}
