package webscraper

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/Fi44er/sdmedik/backend/pkg/webscraper/constants"
	"github.com/Fi44er/sdmedik/backend/pkg/webscraper/structs"
	"github.com/Fi44er/sdmedik/backend/pkg/webscraper/utils"
	"github.com/gofiber/fiber/v2/log"
)

func Scraper(ctx context.Context) []structs.Items {
	regions := constants.Regions
	articles := constants.Articles
	mainUrl := "https://ktsr.sfr.gov.ru"

	log.Info("Starting web scraper")
	doc := utils.Request(mainUrl)
	if doc == nil {
		log.Error("Failed to fetch the main document")
		return nil
	}
	log.Info("Fetched main document from: ", mainUrl)

	sectionsMap := utils.ParseSectionUrl(doc)
	log.Info("Parsed section URLs")

	// Предварительное заполнение articleUrlMap
	articleUrlMap := make(map[string]structs.Category)
	for _, article := range articles {
		articleType := strings.Split(article, "-")[0]
		url := fmt.Sprintf("%v%v", mainUrl, sectionsMap[articleType])
		if _, ok := articleUrlMap[article]; !ok {
			log.Info("Pre-parsing category article URL for article: ", article, " URL: ", url)
			utils.ParseCategoryArticleUrl(url, articleUrlMap)
		}
	}

	// Предварительное заполнение productsMap
	productsMap := make(map[string][]structs.ParseProductsArticlesType)
	for article := range articleUrlMap {
		url := articleUrlMap[article].URL
		if _, ok := productsMap[url]; !ok {
			log.Info("Fetching products articles for URL: ", url)
			productsMap[url] = utils.ParseProductsArticles(url)
		}
	}

	items := make(map[string]structs.Items)
	results := make(chan struct {
		article string
		item    structs.Items
	}, len(articles)*len(regions))

	var wg sync.WaitGroup
	var mu sync.Mutex
	maxGoroutines := 5
	sem := make(chan struct{}, maxGoroutines)

	for _, article := range articles {
		articleType := strings.Split(article, "-")[0]
		log.Info("Processing article: ", article, " with type: ", articleType)
		for _, region := range regions {
			wg.Add(1)
			sem <- struct{}{}

			go func(article, articleType string, region constants.Region) {
				defer wg.Done()
				defer func() { <-sem }()

				select {
				case <-ctx.Done():
					log.Warn("Context done, stopping goroutine for article: ", article, " and region: ", region.Iso3166)
					return
				default:
				}

				certificatePrice := utils.ParceCertificatePriceRegion(region, article, articleType)
				if certificatePrice == nil {
					log.Warn("No certificate price found for article: ", article, " in region: ", region.Iso3166)
					return
				}

				mu.Lock()
				if existingItems, exist := items[article]; exist {
					existingItems.Items = append(existingItems.Items, structs.Item{
						Price:  *certificatePrice,
						Region: region.Iso3166,
					})
					items[article] = existingItems
				} else {
					newItems := structs.Items{
						CategoryArticle: article,
						CategoryName:    articleUrlMap[article].Name,
						Product:         productsMap[articleUrlMap[article].URL],
						Items: []structs.Item{
							{
								Price:  *certificatePrice,
								Region: region.Iso3166,
							},
						},
					}
					items[article] = newItems
				}
				result := items[article]
				mu.Unlock()

				select {
				case <-ctx.Done():
					return
				case results <- struct {
					article string
					item    structs.Items
				}{article, result}:
				}
			}(article, articleType, region)
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		mu.Lock()
		items[result.article] = result.item
		mu.Unlock()
	}

	itemSlice := make([]structs.Items, 0, len(items))
	for _, item := range items {
		itemSlice = append(itemSlice, item)
	}

	log.Info("Scraping completed. Total items scraped: ", len(itemSlice))
	return itemSlice
}
