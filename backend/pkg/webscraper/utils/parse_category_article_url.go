package utils

import (
	"strings"

	"github.com/Fi44er/sdmedik/backend/pkg/webscraper/structs"
	"github.com/PuerkitoBio/goquery"
)

func ParseCategoryArticleUrl(url string, articleUrlMap map[string]structs.Category) {
	doc := Request(url)

	doc.Find("a.category-inner-item-info").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")

		var cartUrl string
		var article string
		if exists {
			s.Find("div.category-inner-item__catalog").Each(func(i int, div *goquery.Selection) {
				text := strings.TrimSpace(div.Text())
				if text != "" {
					article = text
					cartUrl = "https://ktsr.sfr.gov.ru" + href
				}
			})

			s.Find("div.category-inner-item__title").Each(func(i int, div *goquery.Selection) {
				text := strings.TrimSpace(div.Text())
				if text != "" {
					articleWithSpace := article + " "
					name := strings.Replace(text, articleWithSpace, "", 1)
					articleUrlMap[article] = structs.Category{
						URL:  cartUrl,
						Name: name,
					}
				}
			})
		}
	})
}
