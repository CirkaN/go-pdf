package internal

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func generatePdfFromUrl(res *[]byte, url string, waitTime time.Duration) chromedp.Tasks {

	var wg sync.WaitGroup
	wg.Add(1)
	return chromedp.Tasks{
		chromedp.Navigate(url),

		chromedp.ActionFunc(func(ctx context.Context) error {
			lctx, cancel := context.WithCancel(ctx)
			chromedp.ListenTarget(lctx, func(ev interface{}) {
				if _, ok := ev.(*page.EventLoadEventFired); ok {
					wg.Done()
					cancel()
				}
			})
			return nil
		}),
		chromedp.Sleep(waitTime),

		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
func GenerateFromUrl(filePath string, fileName string, url string, waitTime time.Duration) error {
	startTime := time.Now()
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx, generatePdfFromUrl(&buf, url, waitTime)); err != nil {
		return err
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	builtPath := filePath + fileName

	if err := os.WriteFile(builtPath, buf, 0o644); err != nil {
		return err
	}
	fmt.Println("Pdf generated succesfully in /storage path")
	fmt.Printf("PDF generation took %s\n", duration)
	return nil

}
