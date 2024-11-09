package utils

import (
	"context"
	"log"

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
		chromedp.WaitVisible(`.catalog-products__info-price`, chromedp.ByQuery),
		chromedp.Text(`.catalog-products__info-price span`, &price, chromedp.ByQuery),

		// get product article
		chromedp.WaitVisible(`.product-inner-top-article__code`, chromedp.ByQuery),
		chromedp.Text(`.product-inner-top-article__code`, &article, chromedp.ByQuery),
	)
	log.Println(article)
	return price, article, err
}
