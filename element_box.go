package gipdf

type Box struct {
	Ratio        float64
	Width        *float64
	Height       *float64
	Background   Color
	CornerRadius [4]float64 // TL, TR, BR, BL
	Children     []Element
	NoPadding    bool
}

func (b *Box) AspectRatio() float64  { return b.Ratio }
func (b *Box) FixedWidth() *float64  { return b.Width }
func (b *Box) FixedHeight() *float64 { return b.Height }

func (b *Box) Render(ctx *RenderContext, x, y, width, height float64) error {
	ctx.PDF.SetStrokeColor(0, 0, 0)
	ctx.PDF.SetLineWidth(0.2)
	ctx.PDF.SetFillColor(b.Background.R, b.Background.G, b.Background.B)

	// Draw filled rounded rectangle
	err := Rectangle(ctx.PDF, x, y, x+width, y+height, "F", b.CornerRadius, 10)
	if err != nil {
		return err
	}

	ctx.PDF.SetFillColor(0, 0, 0) // Reset fill color for stroke

	if ctx.Debug {
		drawDebugRect(ctx.PDF, x, y, width, height)
	}

	// Render children inside (with optional padding)
	for _, child := range b.Children {
		if !b.NoPadding {
			err := child.Render(ctx, x+max(b.CornerRadius[0]/2, b.CornerRadius[3]/2), y+max(b.CornerRadius[1]/2, b.CornerRadius[2]/2), width-(b.CornerRadius[0]/2)-(b.CornerRadius[1]/2), height-(b.CornerRadius[2]/2)-(b.CornerRadius[3]/2))
			if err != nil {
				return err
			}
		} else {
			err := child.Render(ctx, x, y, width, height)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
