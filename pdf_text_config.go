package gipdf

func UseFont(font Font) configFunc {
	return func(pdf *Document) {
		pdf.SetFont(font.Name, font.Style, font.Size)
	}
}

func UseColor(color Color) configFunc {
	return func(pdf *Document) {
		pdf.SetTextColor(color.Red, color.Green, color.Blue)
	}
}

func UseFontSize(size float64) configFunc {
	return func(pdf *Document) {
		pdf.SetFontSize(size)
	}
}

func UseFillColor(color Color) configFunc {
	return func(pdf *Document) {
		pdf.SetFillColor(color.Red, color.Green, color.Blue)
	}
}
