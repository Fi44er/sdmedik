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

	doc.Find("a.catalog-home-item").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			s.Find("div.catalog-home-item__section").Each(func(i int, div *goquery.Selection) {
				text := strings.TrimSpace(div.Text())
				if text != "" {
					section[removeUnwantedCharacters(text)] = href
				}
			})
		}
	})

	return section
}
