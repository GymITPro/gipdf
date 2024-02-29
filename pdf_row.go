package gipdf

type Row struct {
	Columns     []Widget   `json:"columns"`
	Padding     Padding    `json:"padding"`
	Spacing     float64    `json:"spacing"`
	IsPageBreak bool       `json:"is_page_break"`
	AspectRatio float64    `json:"aspect_ratio"`
	parent      *Column    `json:"-"`
	document    *Document  `json:"-"`
	builder     func(*Row) `json:"-"`
}

func (r *Row) GetAspectRatio() float64 {
	return r.AspectRatio
}

func (r *Row) Render(pdf *Document, x, y, width, height float64) {
	if r.IsPageBreak {
		pdf.AddPage()
		return
	}

	r.Columns = nil
	r.builder(r)

	var columnAspectCount float64
	for _, column := range r.Columns {
		columnAspectCount += column.GetAspectRatio()
	}

	width = width - r.Padding.Left - r.Padding.Right + r.Spacing
	columnUnitWidth := width / columnAspectCount
	xValue := x + r.Padding.Left
	yMax := y
	count := pdf.PageCount()
	for _, column := range r.Columns {
		pdf.SetY(y + r.Padding.Top)
		pdf.SetX(xValue)
		pdf.SetFont(pdf.defaultFont.Name, pdf.defaultFont.Style, pdf.defaultFont.Size)
		column.Render(pdf, pdf.GetX(), pdf.GetY(), columnUnitWidth*column.GetAspectRatio()-r.Spacing, height)
		xValue += columnUnitWidth * column.GetAspectRatio()
		if pdf.GetY() > yMax || pdf.PageCount() > count {
			yMax = pdf.GetY()
			count = pdf.PageCount()
		}
	}
	pdf.SetY(yMax)
	pdf.Ln(r.Padding.Bottom)
}

func (r *Row) Column(padding Padding, spacing, height float64, aspectRatio float64, builder func(*Column)) *Row {
	column := &Column{
		Rows:        nil,
		Padding:     padding,
		Spacing:     spacing,
		AspectRatio: aspectRatio,
		Height:      height,
		parent:      r,
		document:    r.document,
	}
	r.Columns = append(r.Columns, column)
	builder(column)
	return r
}

func (r *Row) Parent() *Column {
	return r.parent
}

func (d *Document) Row(padding Padding, spacing float64, builder func(*Row)) *Document {
	row := &Row{
		Columns:     nil,
		Padding:     padding,
		Spacing:     spacing,
		IsPageBreak: false,
		AspectRatio: 0,
		parent:      nil,
		document:    d,
		builder:     builder,
	}
	d.Rows = append(d.Rows, row)
	return d
}

func (d *Document) PageBreak() *Document {
	row := &Row{
		Columns:     nil,
		Padding:     Padding{},
		Spacing:     0,
		IsPageBreak: true,
		AspectRatio: 0,
		parent:      nil,
		document:    d,
	}
	d.Rows = append(d.Rows, row)
	return d
}
