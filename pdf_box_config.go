package gipdf

import "fmt"

func UseBackgroundColor(color Color) ConfigFunc {
	return func(pdf *Document, x, y, width, height float64, next func()) {
		fmt.Println("UseBackgroundColor", x, y, width, height, color)
		r, g, b := pdf.GetFillColor()
		pdf.SetFillColor(color.Red, color.Green, color.Blue)
		pdf.Rect(x, y, width, height, "F")
		pdf.SetFillColor(r, g, b)
		next()
	}
}
