package gipdf

type textField struct {
	Content     string       `json:"content"`
	AspectRatio float64      `json:"aspect"`
	FixedHeight float64      `json:"fixed_height"`
	FixedWidth  float64      `json:"fixed_width"`
	TextHeight  float64      `json:"text_height"`
	Alignment   string       `json:"alignment"`
	configs     []ConfigFunc `json:"-"`
}

const AlignmentCenter = "C"
const AlignmentLeft = "L"
const AlignmentRight = "R"

func (r *Row) TextField(content string, aspectRatio, textHeight float64, alignment string, configs ...ConfigFunc) *Row {
	r.Columns = append(r.Columns, textField{
		Content:     content,
		AspectRatio: aspectRatio,
		TextHeight:  textHeight,
		Alignment:   alignment,
		configs:     configs,
	})
	return r
}

func (r *Row) TextFieldFixed(content string, width, height, textHeight float64, alignment string, configs ...ConfigFunc) *Row {
	r.Columns = append(r.Columns, textField{
		Content:     content,
		AspectRatio: 0,
		FixedWidth:  width,
		FixedHeight: height,
		TextHeight:  textHeight,
		Alignment:   alignment,
		configs:     configs,
	})
	return r
}

func (r *Column) TextField(content string, aspectRatio, textHeight float64, alignment string, configs ...ConfigFunc) *Column {
	r.Rows = append(r.Rows, textField{
		Content:     content,
		AspectRatio: aspectRatio,
		TextHeight:  textHeight,
		Alignment:   alignment,
		configs:     configs,
	})
	return r
}

func (r *Column) TextFieldFixed(content string, height, width, textHeight float64, alignment string, configs ...ConfigFunc) *Column {
	r.Rows = append(r.Rows, textField{
		Content:     content,
		AspectRatio: 0,
		FixedHeight: height,
		FixedWidth:  width,
		TextHeight:  textHeight,
		Alignment:   alignment,
		configs:     configs,
	})
	return r
}

func (c textField) render(pdf *Document, x, y, width, height float64) {
	configRunner(pdf, x, y, width, height, c.renderI, c.configs...)
}

func (c textField) renderI(pdf *Document, x, y, width, height float64) {
	text := pdf.SplitText(c.Content, width)
	for _, lineText := range text {
		pdf.SetX(x)
		pdf.CellFormat(width, c.TextHeight, lineText, "", 1, c.Alignment, false, 0, "")
	}
}

func (c textField) getAspectRatio() float64 {
	return c.AspectRatio
}

func (c textField) getWidth() float64 {
	return c.FixedWidth
}

func (c textField) getHeight() float64 {
	return c.FixedHeight
}
