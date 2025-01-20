package product

import (
	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// Get godoc
// @Summary Get a product
// @Description Gets a product
// @Tags product
// @Accept json
// @Produce json
// @Param id query string false "Product ID"
// @Param article query string false "Product article"
// @Param name query string false "Product name"
// @Param category_id query integer false "Category ID"
// @Param offset query integer false "Offset"
// @Param limit query integer false "Limit"
// @Param iso query string false "Region ISO"
// @Param minimal query boolean false "Minimal"
// @Param filters query string false "Filters in JSON format" example({"price":{"min":20,"max":100},"characteristics":[{"characteristic_id":1,"values":["string"]}]})
// @Success 200 {object} response.ResponseData "OK"
// @Router /product [get]
func (i *Implementation) Get(ctx *fiber.Ctx) error {
	params := ctx.Queries()
	var criteria dto.ProductSearchCriteria

	utils.BindQueryToStruct(params, &criteria)

	product, err := i.productService.Get(ctx.Context(), criteria)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	if len(*product) == 1 {
		return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": (*product)[0]})
	}
	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": product})
}
