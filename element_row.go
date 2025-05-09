package gipdf

type Row struct {
	Elements []Element
	Ratio    float64 // for nesting
	Width    *float64
	Height   *float64
}

func NewRow(elements ...Element) *Row {
	return &Row{
		Elements: elements,
	}
}

func (r *Row) AspectRatio() float64 {
	return r.Ratio
}

func (r *Row) FixedWidth() *float64 {
	return r.Width
}

func (r *Row) FixedHeight() *float64 {
	return r.Height
}

func (r *Row) Render(ctx *RenderContext, x, y, width, height float64) error {
	totalAspect := 0.0
	totalFixedWidth := 0.0
	rowHeight := 0.0

	// Measure max height of the row
	for _, el := range r.Elements {
		if h := el.FixedHeight(); h != nil && *h > rowHeight {
			rowHeight = *h
		}
	}

	if fh := r.FixedHeight(); fh != nil {
		rowHeight = *fh
	}
	if rowHeight == 0 {
		rowHeight = height
	}

	ctx.EnsureSpace(rowHeight)

	dynamicWidth := width - totalFixedWidth
	for _, el := range r.Elements {
		if w := el.FixedWidth(); w != nil {
			totalFixedWidth += *w
		} else {
			totalAspect += el.AspectRatio()
		}
	}

	currX := x
	for _, el := range r.Elements {
		w := 0.0
		if fw := el.FixedWidth(); fw != nil {
			w = *fw
		} else {
			w = (el.AspectRatio() / totalAspect) * dynamicWidth
		}

		h := rowHeight
		offsetY := ctx.CursorY

		if fh := el.FixedHeight(); fh != nil {
			h = *fh
			offsetY = ctx.CursorY

			// Apply vertical alignment if wrapped in AlignedElement
			if ae, ok := el.(*AlignedElement); ok {
				switch ae.VAlign {
				case "center":
					offsetY = ctx.CursorY + (rowHeight-h)/2
				case "end":
					offsetY = ctx.CursorY + (rowHeight - h)
				}
				el = ae.Element
			}
		}

		if err := el.Render(ctx, currX, offsetY, w, h); err != nil {
			return err
		}
		currX += w
	}

	if ctx.Debug {
		drawDebugRect(ctx.PDF, x, y, width, rowHeight)
	}

	return nil
}
