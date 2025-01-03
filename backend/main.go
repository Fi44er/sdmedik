package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/samber/lo"
)

// Пример использования
func main() {
	data := generateProducts(20)

	val := FilterByPriceRange(data, 10.0, 15.0)

	jsonData, err := json.MarshalIndent(val, "", "  ")
	if err != nil {
		fmt.Println("Ошибка при преобразовании в JSON:", err)
		return
	}

	// Вывод результата
	fmt.Println(string(jsonData))
}

func FilterByCharacteristicID(data []model.Product, characteristicID int) []model.Product {
	return lo.Filter(data, func(product model.Product, _ int) bool {
		return lo.ContainsBy(product.CharacteristicValues, func(charValue model.CharacteristicValue) bool {
			return charValue.CharacteristicID == characteristicID
		})
	})
}

func FilterByPriceRange(data []model.Product, minPrice, maxPrice float64) []model.Product {
	return lo.Filter(data, func(product model.Product, _ int) bool {
		return product.Price >= minPrice && product.Price <= maxPrice
	})
}

// generateProducts создает массив из n товаров
func generateProducts(n int) []model.Product {
	products := make([]model.Product, 0, n)

	for i := 1; i <= n; i++ {
		productID := "product-id-" + strconv.Itoa(i)
		article := "article-" + strconv.Itoa(i)
		name := "Product #" + strconv.Itoa(i)
		description := "Description for product #" + strconv.Itoa(i)
		price := float64(i) * 10.0 // Пример цены

		// Категории
		categories := []model.Category{
			{
				ID:   i,
				Name: "Category #" + strconv.Itoa(i),
			},
		}

		// Изображения
		imageProductID := "image-product-id-" + strconv.Itoa(i)
		images := []model.Image{
			{
				ID:         "image-id-" + strconv.Itoa(i),
				ProductID:  &imageProductID,
				CategoryID: nil,
				Name:       "image-name-" + strconv.Itoa(i) + ".png",
			},
		}

		// Характеристики
		characteristicValues := []model.CharacteristicValue{
			{
				ID:               i,
				Value:            "Value " + strconv.Itoa(i),
				CharacteristicID: i,
				ProductID:        productID,
			},
		}

		// Создание продукта
		product := model.Product{
			ID:                   productID,
			Article:              article,
			Name:                 name,
			Description:          description,
			Price:                price,
			Categories:           categories,
			Certificates:         nil,
			Images:               images,
			CharacteristicValues: characteristicValues,
		}

		// Добавление продукта в массив
		products = append(products, product)
	}

	return products
}
