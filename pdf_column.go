package gipdf

type Column struct {
	Rows        []Widget  `json:"columns"`
	Padding     Padding   `json:"padding"`
	Height      float64   `json:"height"`
	Spacing     float64   `json:"spacing"`
	AspectRatio float64   `json:"aspect_ratio"`
	parent      *Row      `json:"-"`
	document    *Document `json:"-"`
}

func (r *Column) GetAspectRatio() float64 {
	return r.AspectRatio
}

func (r *Column) Render(pdf *Document, x, y, width, height float64) {
	var columnAspectCount float64
	for _, column := range r.Rows {
		columnAspectCount += column.GetAspectRatio()
	}

	tmpHeight := r.Height
	if height > 0 {
		tmpHeight = height
	}
	allHeight := tmpHeight - r.Padding.Top - r.Padding.Bottom + r.Spacing
	columnUnitHeight := allHeight / columnAspectCount
	yValue := y + r.Padding.Top
	xMax := x
	for _, column := range r.Rows {
		pdf.SetY(yValue)
		pdf.SetX(x + r.Padding.Left)
		pdf.SetFont(pdf.defaultFont.Name, pdf.defaultFont.Style, pdf.defaultFont.Size)
		column.Render(pdf, pdf.GetX(), pdf.GetY(), width, columnUnitHeight)
		if pdf.Fpdf.Err() {
			return
		}

		yValue += columnUnitHeight * column.GetAspectRatio()
		yValue += r.Spacing
		if pdf.GetX() > xMax {
			xMax = pdf.GetX()
		}
	}
	pdf.SetX(xMax)
	pdf.SetY(yValue - r.Spacing + r.Padding.Bottom)
}

func (r *Column) Row(padding Padding, spacing float64, aspectRatio float64, builder func(*Row)) *Column {
	row := &Row{
		Columns:     nil,
		Padding:     padding,
		Spacing:     spacing,
		IsPageBreak: false,
		AspectRatio: aspectRatio,
		parent:      r,
		document:    r.document,
		builder:     builder,
	}
	r.Rows = append(r.Rows, row)
	return r
}

func (r *Column) Parent() *Row {
	return r.parent
}

func (r *Column) Done() *Document {
	return r.document
}
