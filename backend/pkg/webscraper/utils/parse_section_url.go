package utils

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseSectionUrl(doc *goquery.Document) map[string]string {
	section := extractLinksAndTexts(doc)
	return section
}

func extractLinksAndTexts(doc *goquery.Document) map[string]string {
	section := make(map[string]string)

	// Ищем все элементы <a> с классом "catalog-home-item"
	doc.Find("a.catalog-home-item").Each(func(i int, s *goquery.Selection) {
		// Извлекаем атрибут href
		href, exists := s.Attr("href")
		if exists {
			// Ищем дочерний элемент <div> с классом "catalog-home-item__section"
			s.Find("div.catalog-home-item__section").Each(func(i int, div *goquery.Selection) {
				// Извлекаем текстовое содержимое
				text := strings.TrimSpace(div.Text())
				if text != "" {
					// Добавляем в мапу, где ключ — текст, значение — ссылка
					section[removeUnwantedCharacters(text)] = href
				}
			})
		}
	})

	return section
}
