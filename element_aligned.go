package gipdf

const (
	AlignStart  = "start"
	AlignCenter = "center"
	AlignEnd    = "end"
)

type AlignedElement struct {
	Element
	HAlign string // "start", "center", "end"
	VAlign string // "start", "center", "end"
}

func Align(el Element, hAlign, vAlign string) *AlignedElement {
	return &AlignedElement{
		Element: el,
		HAlign:  hAlign,
		VAlign:  vAlign,
	}
}
