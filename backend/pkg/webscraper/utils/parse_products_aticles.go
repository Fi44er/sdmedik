package utils

import (
	"fmt"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2/log"
)

func ParseProductsArticles(url string) []string {
	doc := Request(url)
	if doc == nil {
		log.Error("Failed to fetch the initial document")
		return nil
	}

	paginationUrls := parsePaginationUrl(doc)
	if len(paginationUrls) == 0 {
		paginationUrls = []string{url}
	}

	results := make(chan string)
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

	var productsArticles []string
	for article := range results {
		productsArticles = append(productsArticles, article)
	}

	return productsArticles
}

func parseProductsArticlesFromPage(url string, results chan<- string) {
	doc := Request(url)
	if doc == nil {
		log.Errorf("Failed to fetch document from URL: %s", url)
		return
	}

	doc.Find("div.product-item__article").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if text != "" {
			results <- text
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
