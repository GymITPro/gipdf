package gipdf

type dataField struct {
	Title       string       `json:"title"`
	Content     string       `json:"content"`
	AspectRatio float64      `json:"aspect"`
	TitleHeight float64      `json:"title_height"`
	FieldHeight float64      `json:"field_height"`
	configs     []ConfigFunc `json:"-"`
}

func (r *Row) DataField(title, content string, aspectRatio, titleHeight, fieldHeight float64, configs ...ConfigFunc) *Row {
	r.Columns = append(r.Columns, dataField{
		Title:       title,
		Content:     content,
		AspectRatio: aspectRatio,
		TitleHeight: titleHeight,
		FieldHeight: fieldHeight,
		configs:     configs,
	})
	return r
}

func (r *Column) DataField(title, content string, aspectRatio, titleHeight, fieldHeight float64, configs ...ConfigFunc) *Column {
	r.Rows = append(r.Rows, dataField{
		Title:       title,
		Content:     content,
		AspectRatio: aspectRatio,
		TitleHeight: titleHeight,
		FieldHeight: fieldHeight,
		configs:     configs,
	})
	return r
}

func (c dataField) render(pdf *Document, x, y, width, height float64) {
	configRunner(pdf, x, y, width, height, c.renderI, c.configs...)
}

func (c dataField) renderI(pdf *Document, x, y, width, height float64) {
	pdf.CellFormat(width, c.TitleHeight, c.Title, "", 1, "", false, 0, "")

	text := pdf.SplitText(c.Content, width)
	for _, lineText := range text {
		pdf.SetX(x + 1)
		pdf.CellFormat(width, c.FieldHeight, lineText, "", 1, "", true, 0, "")
	}
}

func (c dataField) getAspectRatio() float64 {
	return c.AspectRatio
}
