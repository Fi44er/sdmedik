package utils

import (
	"context"

	"github.com/chromedp/chromedp"
)

func GetProductLinks(ctx context.Context) ([]string, error) {
	var links []string
	err := chromedp.Run(ctx,
		chromedp.Evaluate(`Array.from(document.querySelectorAll('.product-item .product-item-info')).map(a => a.href)`, &links),
	)
	return links, err
}
