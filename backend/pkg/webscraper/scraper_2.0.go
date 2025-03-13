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
	doc := utils.Request(mainUrl)
	sectionsMap := utils.ParseSectionUrl(doc)
	articleUrlMap := make(map[string]structs.Category)
	items := make(map[string]structs.Items)

	results := make(chan struct {
		article string
		item    structs.Items
	})

	var wg sync.WaitGroup
	var mu sync.Mutex

	// Ограничиваем количество одновременно работающих горутин
	maxGoroutines := 20 // Можно настроить в зависимости от возможностей системы
	sem := make(chan struct{}, maxGoroutines)

	for _, article := range articles {
		articleType := strings.Split(article, "-")[0]
		log.Info("article: ", article, " articleType: ", articleType)
		for _, region := range regions {
			wg.Add(1)
			sem <- struct{}{} // Занимаем слот в семафоре

			go func(article string, articleType string, region constants.Region) {
				defer wg.Done()
				defer func() { <-sem }() // Освобождаем слот в семафоре

				select {
				case <-ctx.Done():
					return // Завершаем горутину при отмене контекста
				default:
				}

				certificatePrice := utils.ParceCertificatePriceRegion(region, article, articleType)
				mu.Lock()
				if existingItems, exist := items[article]; exist {
					log.Info("Append price to existing item map")
					existingItems.Items = append(existingItems.Items, structs.Item{
						Price:  *certificatePrice,
						Region: region.Iso3166,
					})
					items[article] = existingItems
					mu.Unlock()
					results <- struct {
						article string
						item    structs.Items
					}{article, existingItems}
				} else {
					log.Info("Create new item map")
					if _, ok := articleUrlMap[article]; !ok {
						url := fmt.Sprintf("%v%v", mainUrl, sectionsMap[articleType])
						utils.ParseCategoryArticleUrl(url, articleUrlMap)
					}

					log.Info(articleUrlMap[article].URL)
					productsArticles := utils.ParseProductsArticles(articleUrlMap[article].URL)
					newItems := structs.Items{
						CategoryArticle: article,
						CategoryName:    articleUrlMap[article].Name,
						Product:         productsArticles,
						Items: []structs.Item{
							{
								Price:  *certificatePrice,
								Region: region.Iso3166,
							},
						},
					}
					items[article] = newItems
					mu.Unlock()
					results <- struct {
						article string
						item    structs.Items
					}{article, newItems}
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

	return itemSlice
}
