package gipdf

type Column struct {
	Elements []Element
	Ratio    float64
	Width    *float64
	Height   *float64
}

func (c *Column) AspectRatio() float64 {
	return c.Ratio
}

func (c *Column) FixedWidth() *float64 {
	return c.Width
}

func (c *Column) FixedHeight() *float64 {
	return c.Height
}

func (c *Column) Render(ctx *RenderContext, x, y, width, height float64) error {
	totalAspect := 0.0
	totalFixedHeight := 0.0

	for _, el := range c.Elements {
		if h := el.FixedHeight(); h != nil {
			totalFixedHeight += *h
		} else {
			totalAspect += el.AspectRatio()
		}
	}

	dynamicHeight := height - totalFixedHeight
	for _, el := range c.Elements {
		h := 0.0
		if fh := el.FixedHeight(); fh != nil {
			h = *fh
		} else {
			h = (el.AspectRatio() / totalAspect) * dynamicHeight
		}

		ctx.EnsureSpace(h)

		w := width
		offsetX := x

		if fw := el.FixedWidth(); fw != nil {
			w = *fw
			offsetX = x

			// Apply horizontal alignment if wrapped in AlignedElement
			if ae, ok := el.(*AlignedElement); ok {
				switch ae.HAlign {
				case "center":
					offsetX = x + (width-w)/2
				case "end":
					offsetX = x + (width - w)
				}
				el = ae.Element
			}
		}

		if err := el.Render(ctx, offsetX, ctx.CursorY, w, h); err != nil {
			return err
		}
		ctx.MoveY(h)
	}

	if ctx.Debug {
		drawDebugRect(ctx.PDF, x, y, width, dynamicHeight)
	}
	return nil
}
