package gipdf

type Padding struct {
	Left   float64 `json:"left"`
	Top    float64 `json:"top"`
	Right  float64 `json:"right"`
	Bottom float64 `json:"bottom"`
}

func PaddingAll(i float64) Padding {
	return Padding{
		Left:   i,
		Top:    i,
		Right:  i,
		Bottom: i,
	}
}

func PaddingLTRB(l, t, r, b float64) Padding {
	return Padding{
		Left:   l,
		Top:    t,
		Right:  r,
		Bottom: b,
	}
}

func PaddingL(i float64) Padding {
	return Padding{
		Left:   i,
		Top:    0,
		Right:  0,
		Bottom: 0,
	}
}

func PaddingR(i float64) Padding {
	return Padding{
		Left:   0,
		Top:    0,
		Right:  i,
		Bottom: 0,
	}
}

func PaddingT(i float64) Padding {
	return Padding{
		Left:   0,
		Top:    i,
		Right:  0,
		Bottom: 0,
	}
}

func PaddingB(i float64) Padding {
	return Padding{
		Left:   0,
		Top:    0,
		Right:  0,
		Bottom: i,
	}
}

type Widget interface {
	Render(pdf *Document, x, y, width, height float64)
	GetAspectRatio() float64
}
