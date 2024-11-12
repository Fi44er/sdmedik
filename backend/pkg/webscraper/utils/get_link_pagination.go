package utils

import (
	"context"

	"github.com/chromedp/chromedp"
)

func GetPaginationLinks(ctx context.Context) ([]string, error) {
	var links []string
	err := chromedp.Run(ctx,
		chromedp.Evaluate(`Array.from(document.querySelectorAll('.pagination .numeric a')).map(a => a.href)`, &links),
	)
	return links, err
}
