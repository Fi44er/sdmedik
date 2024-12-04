package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/Fi44er/sdmedik/backend/pkg/webscraper/structs"
	"github.com/Fi44er/sdmedik/backend/pkg/webscraper/utils"
	"github.com/chromedp/chromedp"
)

const MAIN_URL = "https://ktsr.sfr.gov.ru/ru-RU/product/product/order86n/184"

func main() {
	Run()
}

func Run() {
	var brawserPullLimit = 4

	jobs := make(chan structs.Job, 100)
	results := make(chan structs.Result, 100)

	// var products []structs.Products
	// var productsMu sync.Mutex

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // Отключаем безголовый режим
	)

	allocCtx, cancelAlloc := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancelAlloc()

	ctx, cancelBase := chromedp.NewContext(allocCtx)
	defer cancelBase()

	if err := utils.NavigateTo(ctx, MAIN_URL); err != nil {
		log.Fatal(err)
	}

	pagintioLinks, err := utils.GetPaginationLinks(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var links []string

	if len(pagintioLinks) == 0 {
		pagintioLinks = append(pagintioLinks, MAIN_URL)
	}

	for _, link := range pagintioLinks {
		if err := utils.NavigateTo(ctx, link); err != nil {
			log.Fatal(err)
		}

		productLinks, err := utils.GetProductLinks(ctx)
		if err != nil {
			log.Fatal(err)
		}

		links = append(links, productLinks...)
	}

	go func() {
		for _, link := range links {
			jobs <- structs.Job{Link: link}
		}
		close(jobs) // Закрываем канал после отправки всех заданий
	}()

	var wgWorkers sync.WaitGroup
	if brawserPullLimit > len(links) {
		brawserPullLimit = len(links)
	}

	for i := 0; i < brawserPullLimit; i++ {
		wgWorkers.Add(1)
		go utils.Worker(allocCtx, jobs, results, &wgWorkers)
	}

	products := make(map[string]structs.Products)
	var wgResults sync.WaitGroup
	wgResults.Add(1)
	go func() {
		defer wgResults.Done()
		for res := range results {
			if res.Err != nil {
				log.Printf("Ошибка: %v", res.Err)
				continue
			}

			if product, exist := products[res.Product.Article]; exist {
				product.RegionPrices = append(product.RegionPrices, structs.RegionPrice{
					Region: res.Product.Region,
					Price:  res.Product.Price,
				})
			} else {
				products[res.Product.Article] = structs.Products{
					Article: res.Product.Article,
					RegionPrices: []structs.RegionPrice{
						{
							Region: res.Product.Region,
							Price:  res.Product.Price,
						},
					},
				}
			}
		}
	}()

	wgWorkers.Wait()
	close(results) // Закрываем канал результатов после завершения workers

	wgResults.Wait()

	if err := utils.SaveToFile(products, "products.json"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Данные успешно сохранены в products.json")
}
