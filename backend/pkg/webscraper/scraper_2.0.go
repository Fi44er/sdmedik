package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Fi44er/sdmedik/backend/pkg/webscraper/constants"
	"github.com/Fi44er/sdmedik/backend/pkg/webscraper/structs"
	"github.com/Fi44er/sdmedik/backend/pkg/webscraper/utils"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	regions := constants.Regions
	articles := constants.Articles

	mainUrl := "https://ktsr.sfr.gov.ru"
	doc := utils.Request(mainUrl)

	sectionsMap := utils.ParseSectionUrl(doc)

	articleUrlMap := make(map[string]structs.Category)

	items := make(map[string]structs.Items)

	// Канал для передачи результатов
	results := make(chan struct {
		article string
		item    structs.Items
	})

	// WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Мьютекс для синхронизации доступа к карте items
	var mu sync.Mutex

	// Запускаем горутины для каждой статьи и региона
	for _, article := range articles {
		articleType := strings.Split(article, "-")[0]
		log.Info("article: ", article, " articleType: ", articleType)
		for _, region := range regions {
			wg.Add(1)

			go func(article string, articleType string, region constants.Region) {
				defer wg.Done()
				certificatePrice := utils.ParceCertificatePriceRegion(region, article, articleType)
				mu.Lock()
				if existingItems, exist := items[article]; exist {
					// Если элемент уже существует, добавляем новый регион
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
					// Если элемент не существует, создаем новый
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
						Articles:        productsArticles,
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

	// Горутина для закрытия канала после завершения всех задач
	go func() {
		wg.Wait()
		close(results)
	}()

	// Собираем результаты из канала
	for result := range results {
		mu.Lock()
		items[result.article] = result.item
		mu.Unlock()
	}

	itemSlice := make([]structs.Items, 0, len(items))
	for _, item := range items {
		itemSlice = append(itemSlice, item)
	}

	log.Info("Results: ", itemSlice)
}
