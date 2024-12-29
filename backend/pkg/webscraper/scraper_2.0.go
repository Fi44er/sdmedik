package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Fi44er/sdmedik/backend/pkg/webscraper/constants"
	"github.com/Fi44er/sdmedik/backend/pkg/webscraper/utils"
	"github.com/andybalholm/brotli"
	"golang.org/x/net/html"
)

func main() {
	regions := constants.Regions
	articles := constants.Articles

	type Item struct {
		Price  float64
		Region string
	}

	type Items struct {
		Items []Item
	}

	items := make(map[string]Items)

	for _, article := range articles {
		articleType := strings.Split(article, "-")[0]
		for _, region := range regions {
			url := fmt.Sprintf("https://ktsr.sfr.gov.ru/ru-RU/service/compensation/product-header?region=%v&type=%v&code=%v", region.Iso3166, articleType, article)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				fmt.Println("Ошибка при создании запроса:", err)
				return
			}

			constants.AddHeadersToReq(req)
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Ошибка при выполнении запроса:", err)
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Ошибка при чтении ответа:", err)
				return
			}

			var certificatePrice float64
			if body != nil && resp.StatusCode == 200 {
				br := brotli.NewReader(bytes.NewReader(body))
				decodedBody, err := io.ReadAll(br)
				if err != nil {
					fmt.Println("Ошибка при декодировании Brotli:", err)
					return
				}

				doc, err := html.Parse(strings.NewReader(string(decodedBody)))
				if err != nil {
					fmt.Println("Ошибка при парсинге HTML:", err)
					return
				}

				certificatePrice = utils.ParcePrice(doc)
			} else {
				certificatePrice = 0
			}

			if existingItems, exist := items[article]; exist {
				existingItems.Items = append(items[article].Items, Item{
					Price:  certificatePrice,
					Region: region.Iso3166,
				})
				items[article] = existingItems
			} else {
				newItems := Items{
					Items: []Item{
						{
							Price:  certificatePrice,
							Region: region.Iso3166,
						},
					},
				}
				items[article] = newItems
			}
		}
	}

	fmt.Println(items)
}
