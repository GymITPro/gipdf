package gipdf

type dataField struct {
	Title       string       `json:"title"`
	Content     string       `json:"content"`
	AspectRatio float64      `json:"aspect"`
	TitleHeight float64      `json:"title_height"`
	FieldHeight float64      `json:"field_height"`
	Configs     []configFunc `json:"-"`
}

func (r *Row) DataField(title, content string, aspectRatio, titleHeight, fieldHeight float64, configs ...configFunc) *Row {
	r.Columns = append(r.Columns, dataField{
		Title:       title,
		Content:     content,
		AspectRatio: aspectRatio,
		TitleHeight: titleHeight,
		FieldHeight: fieldHeight,
		Configs:     configs,
	})
	return r
}

func (r *Column) DataField(title, content string, aspectRatio, titleHeight, fieldHeight float64, configs ...configFunc) *Column {
	r.Rows = append(r.Rows, dataField{
		Title:       title,
		Content:     content,
		AspectRatio: aspectRatio,
		TitleHeight: titleHeight,
		FieldHeight: fieldHeight,
		Configs:     configs,
	})
	return r
}

func (c dataField) Render(pdf *Document, x, y, width, height float64) {
	pdf.SetFont("Helvetica", "L", 8)
	for _, config := range c.Configs {
		config(pdf)
	}

	pdf.CellFormat(width, c.TitleHeight, c.Title, "", 1, "", false, 0, "")
	pdf.SetFont("Helvetica", "L", 10)

	text := pdf.SplitText(c.Content, width)
	for _, lineText := range text {
		pdf.SetX(x + 1)
		pdf.CellFormat(width, c.FieldHeight, lineText, "", 1, "", true, 0, "")
	}
}

func (c dataField) GetAspectRatio() float64 {
	return c.AspectRatio
}
