package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/Fi44er/sdmedik/backend/pkg/webscraper/constants"
	"github.com/andybalholm/brotli"
	"golang.org/x/net/html"
)

func ParceCertificatePriceRegion(region constants.Region, article string, articleType string) *float64 {
	url := fmt.Sprintf("https://ktsr.sfr.gov.ru/ru-RU/service/compensation/product-header?region=%v&type=%v&code=%v", region.Iso3166, articleType, article)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return nil
	}

	constants.AddHeadersToReq(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return nil
	}

	var certificatePrice float64
	if body != nil && resp.StatusCode == 200 {
		br := brotli.NewReader(bytes.NewReader(body))
		decodedBody, err := io.ReadAll(br)
		if err != nil {
			fmt.Println("Ошибка при декодировании Brotli:", err)
			return nil
		}

		doc, err := html.Parse(strings.NewReader(string(decodedBody)))
		if err != nil {
			fmt.Println("Ошибка при парсинге HTML:", err)
			return nil
		}

		certificatePrice = ParcePrice(doc)
	} else {
		certificatePrice = 0
	}
	return &certificatePrice
}

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
	cleaned = strings.ReplaceAll(cleaned, ",", ".")

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
