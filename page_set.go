package gipdf

import (
	"github.com/signintech/gopdf"
)

type PageSet struct {
	Pages []*Page
}

func (ps *PageSet) Render(pdf *gopdf.GoPdf, pageSize gopdf.Rect, debug bool) error {
	for _, page := range ps.Pages {
		ctx := newRenderContext(pdf, pageSize.H, page.Margin.Top)
		if debug {
			ctx.Debug = debug
		}
		// Header
		pdf.AddHeader(func() {
			ctx.CursorY = 0
			if page.BackGround != nil {
				pdf.SetFillColor(page.BackGround.R, page.BackGround.G, page.BackGround.B)
				pdf.RectFromUpperLeftWithStyle(0, 0, pageSize.W, pageSize.H, "F")
			}

			if page.Header != nil {
				_ = page.Header.Render(ctx, page.Margin.Left, ctx.CursorY, pageSize.W-page.Margin.Left-page.Margin.Right, page.HeaderHeight())
				ctx.MoveY(page.HeaderHeight())
			}
		})

		// Footer
		if page.Footer != nil {
			pdf.AddFooter(func() {
				pdf.SetY(pageSize.H - page.FooterHeight())
				_ = page.Footer.Render(ctx, page.Margin.Left, pdf.GetY(), pageSize.W-page.Margin.Left-page.Margin.Right, page.FooterHeight())
			})
		} else {
			pdf.AddFooter(nil)
		}

		pdf.AddPage()

		// Content
		err := page.Render(ctx, page.Margin.Left, ctx.CursorY, pageSize.W-page.Margin.Left-page.Margin.Right, pageSize.H)
		if err != nil {
			return err
		}
	}

	return nil
}
