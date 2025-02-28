package core

import (
	"context"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type Chrome struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewChromeInstance() *Chrome {
	return &Chrome{
		//
	}
}

func (ch *Chrome) GetDefaultConfig() []chromedp.ExecAllocatorOption {
	return append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("disable-extensions", true),
	)
}

func (ch *Chrome) LoadConfig(configs []chromedp.ExecAllocatorOption) {
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), configs...)

	browserCtx, cancel := chromedp.NewContext(allocCtx)

	timeoutCtx, cancel := context.WithTimeout(browserCtx, 2*time.Minute)

	ch.ctx = timeoutCtx
	ch.cancel = cancel
}

func (ch *Chrome) CloseContext() {
	ch.cancel()
}

func (ch *Chrome) Print(htmlContent string) ([]byte, error) {
	var res []byte

	return res, chromedp.Run(ch.ctx, chromedp.Tasks{
		chromedp.Navigate("about:blank"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			frameTree, _ := page.GetFrameTree().Do(ctx)

			return page.SetDocumentContent(frameTree.Frame.ID, htmlContent).Do(ctx)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().
				WithPrintBackground(true).
				WithMarginTop(0.4).
				WithMarginBottom(0.4).
				WithMarginLeft(0.4).
				WithMarginRight(0.4).
				Do(ctx)
			res = buf

			return err
		}),
	})
}
