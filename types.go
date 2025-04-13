package gipdf

import "github.com/signintech/gopdf"

type Element interface {
	Render(ctx *RenderContext, x, y, width, height float64) error
	AspectRatio() float64
	FixedWidth() *float64
	FixedHeight() *float64
}

func Ptr[T any](f T) *T {
	return &f
}

type PDF struct {
	pdf     *gopdf.GoPdf
	PageSet *PageSet
	Config  *Config
}

type Margin struct {
	Top, Bottom, Left, Right float64
}

func DefaultMargin() Margin {
	return Margin{Top: 40, Bottom: 40, Left: 40, Right: 40}
}

type Font struct {
	Family  string
	Data    func() ([]byte, error)
	Default bool
}

func drawDebugRect(pdf *gopdf.GoPdf, x, y, width, height float64) {
	pdf.SetLineWidth(0.5)
	pdf.SetStrokeColor(255, 0, 0)
	pdf.RectFromUpperLeftWithStyle(x, y, width, height, "D")
}

type Color struct {
	R, G, B uint8
}

func NewColor(r, g, b uint8) Color {
	return Color{R: r, G: g, B: b}
}
