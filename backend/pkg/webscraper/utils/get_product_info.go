package utils

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func GetPriceAndArticle(ctx context.Context, link, region string) (string, string, error) {
	var price string
	var article string
	err := chromedp.Run(ctx,
		chromedp.Navigate(link),

		// get product price
		chromedp.Evaluate(`document.querySelector(".header__address a.js-open-modal").click();`, nil),
		chromedp.WaitVisible("#modal-city", chromedp.ByID),
		chromedp.SendKeys("#w3", region, chromedp.NodeVisible),
		chromedp.WaitVisible("#suggest-container", chromedp.ByID),
		chromedp.WaitVisible(".ymaps-2-1-79-suggest-item-0", chromedp.ByQuery),
		chromedp.Evaluate(`document.querySelector(".ymaps-2-1-79-suggest-item-0 .ymaps-2-1-79-search__suggest-item").click();`, nil),
		// chromedp.WaitVisible(`.catalog-products__info-price`, chromedp.ByQuery),
		// chromedp.Text(`.catalog-products__info-price span`, &price, chromedp.ByQuery),

		chromedp.ActionFunc(func(ctx context.Context) error {

			// Create a context with a timeout

			timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)

			defer cancel()

			// Wait for the element to be visible

			err := chromedp.WaitVisible(`.catalog-products__info-price`, chromedp.ByQuery).Do(timeoutCtx)

			if err == nil {

				// If the element is visible, get the price

				return chromedp.Text(`.catalog-products__info-price span`, &price, chromedp.ByQuery).Do(ctx)

			}

			// If the element is not visible, just return nil to skip

			return nil

		}),

		// get product article
		chromedp.WaitVisible(`.product-inner-top-article__code`, chromedp.ByQuery),
		chromedp.Text(`.product-inner-top-article__code`, &article, chromedp.ByQuery),
	)
	log.Println(article)
	return price, article, err
}
