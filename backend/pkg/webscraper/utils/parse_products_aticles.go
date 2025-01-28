package utils

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Fi44er/sdmedik/backend/pkg/webscraper/structs"
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2/log"
)

func ParseProductsArticles(url string) []structs.ParseProductsArticlesType {
	doc := Request(url)
	if doc == nil {
		log.Error("Failed to fetch the initial document")
		return nil
	}

	paginationUrls := parsePaginationUrl(doc)
	if len(paginationUrls) == 0 {
		paginationUrls = []string{url}
	}

	results := make(chan structs.ParseProductsArticlesType)
	var wg sync.WaitGroup

	for _, paginationUrl := range paginationUrls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			parseProductsArticlesFromPage(url, results)
		}(paginationUrl)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var productsArticles []structs.ParseProductsArticlesType
	for article := range results {
		productsArticles = append(productsArticles, article)
	}

	return productsArticles
}

func parseProductsArticlesFromPage(url string, results chan<- structs.ParseProductsArticlesType) {
	doc := Request(url)
	if doc == nil {
		log.Errorf("Failed to fetch document from URL: %s", url)
		return
	}

	doc.Find("a.product-item-info").Each(func(i int, s *goquery.Selection) {
		// Извлекаем артикул и название для каждого продукта
		article := strings.TrimSpace(s.Find("div.product-item__article").Text())
		name := strings.TrimSpace(s.Find("div.product-item__title").Text())

		// Если данные найдены, отправляем их в канал
		if article != "" && name != "" {
			results <- structs.ParseProductsArticlesType{
				Article: article,
				Name:    name,
			}
		}
	})
}

func parsePaginationUrl(doc *goquery.Document) []string {
	var paginationUrls []string
	doc.Find("li.numeric a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			paginationUrls = append(paginationUrls, fmt.Sprintf("https://ktsr.sfr.gov.ru%s", href))
		}
	})

	return paginationUrls
}
