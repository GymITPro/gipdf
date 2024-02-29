package gipdf

type emptyField struct {
	AspectRatio float64 `json:"aspect"`
}

func (r *Row) EmptyField(aspectRatio float64) *Row {
	r.Columns = append(r.Columns, emptyField{
		AspectRatio: aspectRatio,
	})
	return r
}

func (r *Column) EmptyField(aspectRatio float64) *Column {
	r.Rows = append(r.Rows, emptyField{
		AspectRatio: aspectRatio,
	})
	return r
}

func (c emptyField) Render(pdf *Document, x, y, width, height float64) {
}

func (c emptyField) GetAspectRatio() float64 {
	return c.AspectRatio
}
