package utils

import (
	"fmt"
	"regexp"
	"strconv"

	"golang.org/x/net/html"
)

func ParcePrice(doc *html.Node) float64 {
	var certificatePrice string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, attr := range n.Attr {
				if attr.Key == "class" && attr.Val == "catalog-products__info-price catalog-products__info-price_space" {
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						if c.Type == html.ElementNode && c.Data == "span" {
							for spanChild := c.FirstChild; spanChild != nil; spanChild = spanChild.NextSibling {
								if spanChild.Type == html.TextNode {
									certificatePrice += spanChild.Data
								}
							}
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	cleaned := removeUnwantedCharacters(certificatePrice)

	value, err := strconv.ParseFloat(cleaned, 64)
	if err != nil {
		fmt.Println("Ошибка при парсинге:", err)
		return 0
	}

	return value
}

func removeUnwantedCharacters(s string) string {
	re := regexp.MustCompile("[^0-9.,]+")
	return re.ReplaceAllString(s, "")
}
