package gipdf

import "fmt"

func UseFont(font Font) ConfigFunc {
	return func(pdf *Document, x, y, width, height float64, next func()) {
		fmt.Println("UseFont", font)
		pdf.SetFont(font.Name, font.Style, font.Size)
		next()
		pdf.SetFont(pdf.defaultFont.Name, pdf.defaultFont.Style, pdf.defaultFont.Size)
	}
}

func UseColor(color Color) ConfigFunc {
	return func(pdf *Document, x, y, width, height float64, next func()) {
		fmt.Println("UseColor", color)
		r, g, b := pdf.GetTextColor()
		pdf.SetTextColor(color.Red, color.Green, color.Blue)
		next()
		pdf.SetTextColor(r, g, b)
	}
}

func UseFontSize(size float64) ConfigFunc {
	return func(pdf *Document, x, y, width, height float64, next func()) {
		pdf.SetFontSize(size)
		next()
		pdf.SetFontSize(pdf.defaultFont.Size)
	}
}

func UseFillColor(color Color) ConfigFunc {
	return func(pdf *Document, x, y, width, height float64, next func()) {
		r, g, b := pdf.GetFillColor()
		pdf.SetFillColor(color.Red, color.Green, color.Blue)
		next()
		pdf.SetFillColor(r, g, b)
	}
}
