package gipdf

type Row struct {
	Columns     []Widget     `json:"columns"`
	Padding     Padding      `json:"padding"`
	Spacing     float64      `json:"spacing"`
	IsPageBreak bool         `json:"is_page_break"`
	AspectRatio float64      `json:"aspect_ratio"`
	FixedHeight float64      `json:"fixed_height"`
	FixedWidth  float64      `json:"fixed_width"`
	builder     func(*Row)   `json:"-"`
	configs     []ConfigFunc `json:"-"`
}

func (r *Row) getAspectRatio() float64 {
	return r.AspectRatio
}

func (r *Row) getWidth() float64 {
	return r.FixedWidth
}

func (r *Row) getHeight() float64 {
	return r.FixedHeight
}

func (r *Row) render(pdf *Document, x, y, width, height float64) {
	if r.IsPageBreak {
		pdf.AddPage()
		return
	}

	configRunner(pdf, x, y, width, height, r.renderI, r.configs...)
}

func (r *Row) renderI(pdf *Document, x, y, width, height float64) {
	if r.IsPageBreak {
		pdf.AddPage()
		return
	}

	r.Columns = nil
	r.builder(r)

	var columnAspectCount float64
	var reservedWidth float64
	var maxFixedHeight float64
	for _, column := range r.Columns {
		if a, ok := isAspectRatio(column); ok {
			columnAspectCount += a.getAspectRatio()
		}

		if a, ok := isFixedWidth(column); ok {
			reservedWidth += a.getWidth()
		}

		if a, ok := isFixedHeight(column); ok {
			maxFixedHeight = max(maxFixedHeight, a.getHeight())
		}
	}

	// No negative height
	maxFixedHeight = max(maxFixedHeight, 0) + y

	width = width - r.Padding.Left - r.Padding.Right - (r.Spacing * float64(len(r.Columns)-1)) - reservedWidth
	columnUnitWidth := width / columnAspectCount
	xValue := x + r.Padding.Left
	yMax := y
	count := pdf.PageCount()
	for _, column := range r.Columns {
		pdf.SetY(y + r.Padding.Top)
		pdf.SetX(xValue)
		pdf.SetFont(pdf.defaultFont.Name, pdf.defaultFont.Style, pdf.defaultFont.Size)
		var width float64
		if a, ok := isAspectRatio(column); ok {
			width = max(width, columnUnitWidth*a.getAspectRatio())
		}

		if a, ok := isFixedWidth(column); ok {
			width = max(width, a.getWidth())
		}

		if a, ok := isFixedHeight(column); ok {
			overrideHeight := a.getHeight()
			if overrideHeight > 0 {
				height = overrideHeight
			}
		}

		column.render(pdf, pdf.GetX(), pdf.GetY(), width, height)
		if pdf.Fpdf.Err() {
			return
		}

		xValue += r.Spacing
		if a, ok := isAspectRatio(column); ok && a.getAspectRatio() > 0 {
			xValue += columnUnitWidth * a.getAspectRatio()
		}

		if a, ok := isFixedWidth(column); ok && a.getWidth() > 0 {
			xValue += a.getWidth()
		}

		if pdf.GetY() > yMax || pdf.PageCount() > count {
			yMax = pdf.GetY()
			count = pdf.PageCount()
		}
	}
	pdf.SetY(max(yMax, maxFixedHeight))
	pdf.Ln(r.Padding.Bottom)
}

func (r *Row) Column(padding Padding, spacing, minHeight float64, aspectRatio float64, builder func(*Column), configs ...ConfigFunc) *Row {
	column := &Column{
		Rows:        nil,
		Padding:     padding,
		Spacing:     spacing,
		AspectRatio: aspectRatio,
		MinHeight:   minHeight,
		configs:     configs,
	}
	r.Columns = append(r.Columns, column)
	builder(column)
	return r
}

func (r *Row) ColumnFixed(padding Padding, spacing, height, width float64, aspectRatio float64, builder func(*Column), configs ...ConfigFunc) *Row {
	column := &Column{
		Rows:        nil,
		Padding:     padding,
		Spacing:     spacing,
		AspectRatio: aspectRatio,
		FixedHeight: height,
		FixedWidth:  width,
		configs:     configs,
	}
	r.Columns = append(r.Columns, column)
	builder(column)
	return r
}

func (d *Document) Row(padding Padding, spacing float64, builder func(*Row), configs ...ConfigFunc) *Document {
	row := &Row{
		Columns:     nil,
		Padding:     padding,
		Spacing:     spacing,
		IsPageBreak: false,
		AspectRatio: 0,
		builder:     builder,
		configs:     configs,
	}
	d.Rows = append(d.Rows, row)
	return d
}

func (d *Document) RowFixed(padding Padding, spacing float64, height, width float64, builder func(*Row), configs ...ConfigFunc) *Document {
	row := &Row{
		Columns:     nil,
		Padding:     padding,
		Spacing:     spacing,
		FixedHeight: height,
		FixedWidth:  width,
		IsPageBreak: false,
		AspectRatio: 0,
		builder:     builder,
		configs:     configs,
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
	}
	d.Rows = append(d.Rows, row)
	return d
}
