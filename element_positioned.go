package gipdf

// PositionedElement represents an element that is positioned at a specific location on the page.
// Can not be used with an aspect ratio!
type PositionedElement struct {
	X, Y    float64  // Absolute position on the page (in points)
	Width   *float64 // Optional fixed width
	Height  *float64 // Optional fixed height
	Element Element  // Wrapped inner element
}

func (p *PositionedElement) AspectRatio() float64 {
	return 0 // ignored
}

func (p *PositionedElement) FixedWidth() *float64 {
	return Ptr(0.0) // ignored
}

func (p *PositionedElement) FixedHeight() *float64 {
	return Ptr(0.0) // ignored
}

func (p *PositionedElement) Render(ctx *RenderContext, x, y, _, _ float64) error {
	width := 0.0
	height := 0.0

	if p.Width != nil {
		width = *p.Width
	}
	if p.Height != nil {
		height = *p.Height
	}

	err := p.Element.Render(ctx, p.X, p.Y, width, height)
	if err != nil {
		return err
	}

	if ctx.Debug {
		drawDebugRect(ctx.PDF, p.X, p.Y, width, height)
	}

	// Reset Position
	ctx.PDF.SetXY(x, y)
	return nil
}
