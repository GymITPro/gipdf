package gipdf

import "github.com/signintech/gopdf"

type Image struct {
	Path   string
	Ratio  float64
	Width  *float64
	Height *float64
}

func (i *Image) AspectRatio() float64  { return i.Ratio }
func (i *Image) FixedWidth() *float64  { return i.Width }
func (i *Image) FixedHeight() *float64 { return i.Height }

func (i *Image) Render(ctx *RenderContext, x, y, width, height float64) error {
	err := ctx.PDF.Image(i.Path, x, y, &gopdf.Rect{W: width, H: height})
	if err != nil {
		return err
	}

	if ctx.Debug {
		drawDebugRect(ctx.PDF, x, y, width, height)
	}

	return nil
}
