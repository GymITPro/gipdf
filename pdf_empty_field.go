package gipdf

type emptyField struct {
	AspectRatio float64      `json:"aspect"`
	configs     []ConfigFunc `json:"-"`
}

func (r *Row) EmptyField(aspectRatio float64, configs ...ConfigFunc) *Row {
	r.Columns = append(r.Columns, emptyField{
		AspectRatio: aspectRatio,
		configs:     configs,
	})
	return r
}

func (r *Column) EmptyField(aspectRatio float64, configs ...ConfigFunc) *Column {
	r.Rows = append(r.Rows, emptyField{
		AspectRatio: aspectRatio,
		configs:     configs,
	})
	return r
}

func (c emptyField) render(pdf *Document, x, y, width, height float64) {
	configRunner(pdf, x, y, width, height, func(pdf *Document, x, y, width, height float64) {
	}, c.configs...)
}

func (c emptyField) getAspectRatio() float64 {
	return c.AspectRatio
}
