package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/Fi44er/sdmedik/backend/pkg/webscraper/constants"
)

type Item struct {
	Price  float64 `json:"price"`
	Region string  `json:"region"`
}

type ParseProductsArticlesType struct {
	Article string `json:"article"`
	Name    string `json:"name"`
}

type Items struct {
	CategoryArticle string                      `json:"category_article"`
	CategoryName    string                      `json:"category_name"`
	Items           []Item                      `json:"items"`
	Product         []ParseProductsArticlesType `json:"product"`
}

func parseItems(data string) ([]Items, error) {
	var result []Items

	// Регулярное выражение для категорий:
	// Предполагаем, что код категории - это XX-XX-XX, за которым идет пробел и название
	categoryPattern := regexp.MustCompile(`\{(\d{2}-\d{2}-\d{2})\s+([^}]+?)\s+(\[\{.*?\}\])\s+(\[\{.*?\}\])\}`)
	matches := categoryPattern.FindAllStringSubmatch(data, -1)

	for _, match := range matches {
		if len(match) != 5 {
			continue
		}

		categoryArticle := match[1] // Код категории (например, "17-01-16")
		categoryName := match[2]    // Название категории (все после кода до списка Items)
		itemsStr := match[3]        // Список цен и регионов
		productsStr := match[4]     // Список продуктов

		// Парсинг Items
		var items []Item
		itemsPattern := regexp.MustCompile(`\{([\d.]+)\s+([A-Za-z0-9-]+)\}`)
		itemMatches := itemsPattern.FindAllStringSubmatch(itemsStr, -1)
		for _, itemMatch := range itemMatches {
			if len(itemMatch) != 3 {
				continue
			}
			var price float64
			fmt.Sscanf(itemMatch[1], "%f", &price)
			items = append(items, Item{
				Price:  price,
				Region: itemMatch[2],
			})
		}

		// Парсинг Product
		var products []ParseProductsArticlesType
		productsPattern := regexp.MustCompile(`\{([0-9-.]+)\s+([^}]+)\}`)
		productMatches := productsPattern.FindAllStringSubmatch(productsStr, -1)
		for _, prodMatch := range productMatches {
			if len(prodMatch) != 3 {
				continue
			}
			products = append(products, ParseProductsArticlesType{
				Article: prodMatch[1],
				Name:    strings.TrimSpace(prodMatch[2]),
			})
		}

		// Собираем структуру
		result = append(result, Items{
			CategoryArticle: categoryArticle,
			CategoryName:    strings.TrimSpace(categoryName),
			Items:           items,
			Product:         products,
		})
	}

	return result, nil
}

func main() {
	// Чтение файла
	inputFile := "logs.txt"
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Ошибка чтения файла: %v\n", err)
		return
	}

	// Парсинг
	items, err := parseItems(string(data))
	if err != nil {
		fmt.Printf("Ошибка парсинга: %v\n", err)
		return
	}

	// Преобразование в JSON
	jsonData, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		fmt.Printf("Ошибка преобразования в JSON: %v\n", err)
		return
	}

	// Сохранение в файл
	outputFile := "output.json"
	err = os.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		fmt.Printf("Ошибка записи в файл: %v\n", err)
		return
	}

	fmt.Println("Парсинг завершен, результат сохранен в output.json")

	// Создаем map для быстрого поиска CategoryArticle
	foundArticles := make(map[string]bool)
	for _, item := range items {
		foundArticles[item.CategoryArticle] = true
	}

	// Проверяем, каких артикулов нет в результате парсинга
	missingArticles := []string{}
	for _, article := range constants.Articles {
		if !foundArticles[article] {
			missingArticles = append(missingArticles, article)
		}
	}

	// Выводим отсутствующие артикулы
	if len(missingArticles) > 0 {
		fmt.Println("Следующие артикулы отсутствуют в результате парсинга:")
		for _, missing := range missingArticles {
			fmt.Println(missing)
		}
	} else {
		fmt.Println("Все артикулы из массива Articles найдены в результате парсинга.")
	}
}
