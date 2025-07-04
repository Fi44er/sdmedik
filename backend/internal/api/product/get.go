package product

import (
	"strconv"
	"strings"

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
// @Param monotony_price query boolean false "MonotonyPrice"
// @Param catalogs query []int false "Catalogs (comma-separated)" collectionFormat(csv)
// @Param filters query string false "Filters in JSON format" example({"price":{"min":20,"max":100},"characteristics":[{"characteristic_id":1,"values":["string"]}]})
// @Success 200 {object} response.ResponseData "OK"
// @Router /product [get]
func (i *Implementation) Get(ctx *fiber.Ctx) error {
	params := ctx.Queries()
	var criteria dto.ProductSearchCriteria

	if catalogsStr := params["catalogs"]; catalogsStr != "" {
		catalogs, err := stringToIntSlice(catalogsStr)
		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": "invalid catalogs format",
			})
		}
		criteria.Catalogs = catalogs
	}

	utils.BindQueryToStruct(params, &criteria)

	product, count, err := i.productService.Get(ctx.Context(), criteria)
	if err != nil {
		code, msg := errors.GetErroField(err)
		return ctx.Status(code).JSON(msg)
	}

	if len(*product) == 1 {
		return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": (*product)[0]})
	}
	return ctx.Status(200).JSON(fiber.Map{"status": "success", "data": product, "count": count})
}

func stringToIntSlice(s string) ([]int, error) {
	if s == "" {
		return nil, nil
	}

	parts := strings.Split(s, ",")
	result := make([]int, 0, len(parts))

	for _, part := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}

	return result, nil
}
