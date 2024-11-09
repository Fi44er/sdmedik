package utils

import (
	"context"

	"github.com/chromedp/chromedp"
)

func NavigateTo(ctx context.Context, url string) error {
	return chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
	)
}
