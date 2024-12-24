package utils

import (
	"context"
	"log"
	"sync"

	"github.com/Fi44er/sdmedik/backend/pkg/webscraper/constants"
	"github.com/Fi44er/sdmedik/backend/pkg/webscraper/structs"
	"github.com/chromedp/chromedp"
)

func Worker(baseCtx context.Context, jobs <-chan structs.Job, results chan<- structs.Result, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, cancel := chromedp.NewContext(baseCtx)
	defer cancel()

	for job := range jobs {
		for _, region := range constants.Subjects {
			log.Println("region: ", region)
			price, article, err := GetPriceAndArticle(ctx, job.Link, region)
			if err != nil {
				log.Println(err)
				results <- structs.Result{Err: err}
				continue
			}

			product := structs.Product{
				Region:  region,
				Price:   price,
				Article: article,
			}

			results <- structs.Result{Product: product}
		}
	}
}
