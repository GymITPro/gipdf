package gipdf

import (
	"fmt"

	"github.com/signintech/gopdf"
)

type RenderContext struct {
	PDF        *gopdf.GoPdf
	PageHeight float64
	PageWidth  float64
	CursorY    float64
	Debug      bool
}

func newRenderContext(pdf *gopdf.GoPdf, pageHeight, pageWidth float64) *RenderContext {
	return &RenderContext{
		PDF:        pdf,
		PageHeight: pageHeight,
		PageWidth:  pageWidth,
		CursorY:    0,
		Debug:      false,
	}
}

func (ctx *RenderContext) EnsureSpace(heightNeeded float64) {
	fmt.Println("EnsureSpace", ctx.CursorY, heightNeeded, ctx.PageHeight)
	if ctx.CursorY+heightNeeded > ctx.PageHeight {
		ctx.PDF.AddPage()
		ctx.CursorY = 0
	}
}

func (ctx *RenderContext) MoveY(offset float64) {
	ctx.CursorY += offset
}
