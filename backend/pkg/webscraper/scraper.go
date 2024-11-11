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

const BROWSER_PULL_LIMIT = 2
const MAIN_URL = "https://ktsr.sfr.gov.ru/ru-RU/product/product/order86n/81"

func main() {
	Run()
}

func Run() {

	jobs := make(chan structs.Job, 100)
	results := make(chan structs.Result, 100)

	var products []structs.Product
	var productsMu sync.Mutex

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

	links, err := utils.GetProductLinks(ctx)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for _, link := range links {
			jobs <- structs.Job{Link: link}
		}
		close(jobs) // Закрываем канал после отправки всех заданий
	}()

	var wgWorkers sync.WaitGroup
	if BROWSER_PULL_LIMIT > len(links) {
		for i := 0; i < len(links); i++ {
			wgWorkers.Add(1)
			go utils.Worker(allocCtx, jobs, results, &wgWorkers)
		}
	} else {
		for i := 0; i < BROWSER_PULL_LIMIT; i++ {
			wgWorkers.Add(1)
			go utils.Worker(allocCtx, jobs, results, &wgWorkers)
		}
	}

	var wgResults sync.WaitGroup
	wgResults.Add(1)
	go func() {
		defer wgResults.Done()
		for res := range results {
			if res.Err != nil {
				log.Printf("Ошибка: %v", res.Err)
				continue
			}
			productsMu.Lock()
			products = append(products, res.Product)
			productsMu.Unlock()
		}
	}()

	wgWorkers.Wait()
	close(results) // Закрываем канал результатов после завершения workers

	// Ожидаем завершения сбора результатов
	wgResults.Wait()

	// Сохраняем результаты в файл
	if err := utils.SaveToFile(products, "products.json"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Данные успешно сохранены в products.json")
}
