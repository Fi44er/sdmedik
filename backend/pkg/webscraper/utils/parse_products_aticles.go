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
	log.Info("Fetching products articles from URL: ", url)
	doc := Request(url)
	if doc == nil {
		log.Error("Failed to fetch the initial document from URL: ", url)
		return nil
	}

	paginationUrls := parsePaginationUrl(doc)
	if len(paginationUrls) == 0 {
		log.Warn("No pagination URLs found, using the initial URL for parsing")
		paginationUrls = []string{url}
	} else {
		log.Info("Found pagination URLs: ", paginationUrls)
	}

	results := make(chan structs.ParseProductsArticlesType)
	var wg sync.WaitGroup

	maxGoroutines := 5
	semaphore := make(chan struct{}, maxGoroutines)

	for _, paginationUrl := range paginationUrls {
		wg.Add(1)
		semaphore <- struct{}{} // Занимаем слот
		go func(url string) {
			defer wg.Done()
			defer func() { <-semaphore }() // Освобождаем слот
			log.Info("Parsing products from pagination URL: ", url)
			parseProductsArticlesFromPage(url, results)
		}(paginationUrl)
	}

	go func() {
		wg.Wait()
		close(results)
		log.Info("All pagination URLs have been processed, closing results channel")
	}()

	var productsArticles []structs.ParseProductsArticlesType
	for article := range results {
		productsArticles = append(productsArticles, article)
		log.Info("Added product article: ", article.Article, " with name: ", article.Name)
	}

	log.Info("Total products articles parsed: ", len(productsArticles))
	return productsArticles
}

func parseProductsArticlesFromPage(url string, results chan<- structs.ParseProductsArticlesType) {
	log.Info("Fetching document from URL: ", url)
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
			log.Info("Extracted product - Article: ", article, ", Name: ", name)
		} else {
			log.Warn("Missing article or name for product at index: ", i)
		}
	})
}

func parsePaginationUrl(doc *goquery.Document) []string {
	var paginationUrls []string
	doc.Find("li.numeric a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			paginationUrls = append(paginationUrls, fmt.Sprintf("https://ktsr.sfr.gov.ru%s", href))
			log.Info("Found pagination URL: ", href)
		} else {
			log.Warn("No href attribute found for pagination link at index: ", i)
		}
	})

	return paginationUrls
}
