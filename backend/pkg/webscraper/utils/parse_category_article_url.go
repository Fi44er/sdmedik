package utils

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseCategoryArticleUrl(url string, articleUrlMap map[string]string) {
	doc := Request(url)

	doc.Find("a.category-inner-item-info").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		// log.Fatal("БЛЯЯЯЯЯЯть")

		if exists {
			s.Find("div.category-inner-item__catalog").Each(func(i int, div *goquery.Selection) {
				text := strings.TrimSpace(div.Text())
				if text != "" {
					articleUrlMap[text] = "https://ktsr.sfr.gov.ru" + href
				}
			})
		}
	})

}
