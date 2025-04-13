package gipdf

import "github.com/signintech/gopdf"

type Text struct {
	Text     string
	Ratio    float64
	Width    *float64
	Height   *float64
	FontSize float64
	FontName string
	Color    *Color
}

func (t *Text) AspectRatio() float64 {
	return t.Ratio
}

func (t *Text) FixedWidth() *float64 {
	return t.Width
}

func (t *Text) FixedHeight() *float64 {
	return t.Height
}

func (t *Text) Render(ctx *RenderContext, x, y, width, height float64) error {
	if t.Color != nil {
		ctx.PDF.SetTextColor(t.Color.R, t.Color.G, t.Color.B)
	} else {
		ctx.PDF.SetTextColor(0, 0, 0)
	}

	err := ctx.PDF.SetFont(t.FontName, "", t.FontSize)
	if err != nil {
		return err
	}

	if t.Height == nil {
		t.Height = Ptr(t.FontSize * 1.5)
	}

	ctx.PDF.SetX(x)
	ctx.PDF.SetY(y)
	err = ctx.PDF.CellWithOption(&gopdf.Rect{W: width, H: *t.Height}, t.Text, gopdf.CellOption{Align: gopdf.Left})
	if err != nil {
		return err
	}

	ctx.PDF.SetFillColor(0, 0, 0) // Reset fill color for stroke

	if ctx.Debug {
		drawDebugRect(ctx.PDF, x, y, width, *t.Height)
	}

	return nil
}
