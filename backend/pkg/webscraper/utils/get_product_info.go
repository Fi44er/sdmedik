package utils

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
)

func GetPriceAndArticle(ctx context.Context, link, region string) (string, string, error) {
	var price string
	var article string
	err := chromedp.Run(ctx,
		chromedp.Navigate(link),

		chromedp.Evaluate(`document.querySelector(".header__address a.js-open-modal").click();`, nil),
		chromedp.WaitVisible("#modal-city", chromedp.ByID),
		chromedp.SendKeys("#w1", region, chromedp.NodeVisible),
		chromedp.WaitVisible("#suggest-container", chromedp.ByID),
		chromedp.WaitVisible(".ymaps-2-1-79-suggest-item-0", chromedp.ByQuery),
		chromedp.Evaluate(`document.querySelector(".ymaps-2-1-79-suggest-item-0 .ymaps-2-1-79-search__suggest-item").click();`, nil),

		chromedp.ActionFunc(func(ctx context.Context) error {
			timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()
			err := chromedp.WaitVisible(`.catalog-products__info-price`, chromedp.ByQuery).Do(timeoutCtx)
			if err == nil {
				return chromedp.Text(`.catalog-products__info-price span`, &price, chromedp.ByQuery).Do(ctx)
			}
			return nil
		}),
		chromedp.WaitVisible(`.product-inner-top-article__code`, chromedp.ByQuery),
		chromedp.Text(`.product-inner-top-article__code`, &article, chromedp.ByQuery),
	)
	return price, article, err
}
