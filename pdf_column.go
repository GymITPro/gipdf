package gipdf

type Column struct {
	Rows        []Widget     `json:"columns"`
	Padding     Padding      `json:"padding"`
	Height      float64      `json:"height"`
	Spacing     float64      `json:"spacing"`
	AspectRatio float64      `json:"aspect_ratio"`
	configs     []ConfigFunc `json:"-"`
}

func (r *Column) getAspectRatio() float64 {
	return r.AspectRatio
}

func (r *Column) render(pdf *Document, x, y, width, height float64) {
	configRunner(pdf, x, y, width, height, r.renderI, r.configs...)
}

func (r *Column) renderI(pdf *Document, x, y, width, height float64) {
	var columnAspectCount float64
	var maxFixedWidth float64
	var reservedHeight float64
	for _, row := range r.Rows {
		if a, ok := isAspectRatio(row); ok {
			columnAspectCount += a.getAspectRatio()
		}

		if a, ok := isFixedHeight(row); ok {
			reservedHeight += a.getHeight()
		}

		if a, ok := isFixedWidth(row); ok {
			maxFixedWidth = max(maxFixedWidth, a.getWidth())
		}
	}

	// No negative height
	maxFixedWidth = max(maxFixedWidth, 0) + x

	tmpHeight := r.Height
	if height > 0 {
		tmpHeight = height
	}
	allHeight := tmpHeight - r.Padding.Top - r.Padding.Bottom - (r.Spacing * float64(len(r.Rows)-1)) - reservedHeight
	columnUnitHeight := allHeight / columnAspectCount
	yValue := y + r.Padding.Top
	xMax := x
	for _, row := range r.Rows {
		pdf.SetY(yValue)
		pdf.SetX(x + r.Padding.Left)
		pdf.SetFont(pdf.defaultFont.Name, pdf.defaultFont.Style, pdf.defaultFont.Size)
		var height float64
		if a, ok := isAspectRatio(row); ok {
			height = max(height, columnUnitHeight*a.getAspectRatio())
		}

		if a, ok := isFixedHeight(row); ok {
			height = max(height, a.getHeight())
		}

		if a, ok := isFixedWidth(row); ok {
			overrideWidth := a.getWidth()
			if overrideWidth > 0 {
				width = overrideWidth
			}
		}

		row.render(pdf, pdf.GetX(), pdf.GetY(), width, height)
		if pdf.Fpdf.Err() {
			return
		}

		yValue += r.Spacing
		if a, ok := isAspectRatio(row); ok {
			yValue += columnUnitHeight * a.getAspectRatio()
		}

		if a, ok := isFixedHeight(row); ok {
			yValue += a.getHeight()
		}

		if pdf.GetX() > xMax {
			xMax = pdf.GetX()
		}
	}
	pdf.SetX(max(xMax, maxFixedWidth))
	pdf.SetY(yValue - r.Spacing + r.Padding.Bottom)
}

func (r *Column) Row(padding Padding, spacing float64, aspectRatio float64, builder func(*Row), config ...ConfigFunc) *Column {
	row := &Row{
		Columns:     nil,
		Padding:     padding,
		Spacing:     spacing,
		IsPageBreak: false,
		AspectRatio: aspectRatio,
		builder:     builder,
		configs:     config,
	}
	r.Rows = append(r.Rows, row)
	return r
}
