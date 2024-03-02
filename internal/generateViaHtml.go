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


func printToPDF(res *[]byte, html string, waitTime time.Duration) chromedp.Tasks {
	var wg sync.WaitGroup
	wg.Add(1)
	return chromedp.Tasks{
		chromedp.Navigate("about:blank"),
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
		chromedp.ActionFunc(func(ctx context.Context) error {
			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}
			return page.SetDocumentContent(frameTree.Frame.ID, html).Do(ctx)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			wg.Wait()
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
func GenerateFromHtml(filePath string, fileName string, html string, waitTime time.Duration) error {
	startTime := time.Now()
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()

	htmlContent, err := os.ReadFile(html)
	if err != nil {
		return err
	}

	var buf []byte
	if err := chromedp.Run(ctx, printToPDF(&buf, string(htmlContent), waitTime)); err != nil {
		return err;
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	builtPath := filePath + fileName

	if err := os.WriteFile(builtPath, buf, 0o644); err != nil {
		return err;
	}
	fmt.Println("Pdf generated succesfully in /storage path")
	fmt.Printf("PDF generation took %s\n", duration)
	return nil;
}