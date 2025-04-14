package gipdf

import (
	"fmt"

	"github.com/signintech/gopdf"
)

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

func HexColor(hex string) Color {
	if len(hex) != 7 {
		panic("hex color must be in the format #RRGGBB")
	}

	r := uint8(0)
	g := uint8(0)
	b := uint8(0)

	_, err := fmt.Sscanf(hex, "#%02X%02X%02X", &r, &g, &b)
	if err != nil {
		panic(err)
	}

	return Color{R: r, G: g, B: b}
}
