package gipdf

type textField struct {
	Content     string       `json:"content"`
	AspectRatio float64      `json:"aspect"`
	TextHeight  float64      `json:"text_height"`
	Alignment   string       `json:"alignment"`
	Configs     []configFunc `json:"-"`
}

const AlignmentCenter = "C"
const AlignmentLeft = "L"
const AlignmentRight = "R"

func (r *Row) TextField(content string, aspectRatio, textHeight float64, alignment string, configs ...configFunc) *Row {
	r.Columns = append(r.Columns, textField{
		Content:     content,
		AspectRatio: aspectRatio,
		TextHeight:  textHeight,
		Alignment:   alignment,
		Configs:     configs,
	})
	return r
}

func (r *Column) TextField(content string, aspectRatio, textHeight float64, alignment string, configs ...configFunc) *Column {
	r.Rows = append(r.Rows, textField{
		Content:     content,
		AspectRatio: aspectRatio,
		TextHeight:  textHeight,
		Alignment:   alignment,
		Configs:     configs,
	})
	return r
}

func (c textField) Render(pdf *Document, x, y, width, height float64) {
	for _, config := range c.Configs {
		config(pdf)
	}

	text := pdf.SplitText(c.Content, width)
	for _, lineText := range text {
		pdf.SetX(x)
		pdf.CellFormat(width, c.TextHeight, lineText, "", 1, c.Alignment, false, 0, "")
	}
}

func (c textField) GetAspectRatio() float64 {
	return c.AspectRatio
}
