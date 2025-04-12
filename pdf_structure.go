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
	render(pdf *Document, x, y, width, height float64)
}

type aspectRatio interface {
	getAspectRatio() float64
}

func isAspectRatio(w interface{}) (aspectRatio, bool) {
	a, ok := w.(aspectRatio)
	if !ok || a.getAspectRatio() <= 0 {
		return nil, false
	}
	return a, ok
}

type fixedWidth interface {
	getWidth() float64
}

func isFixedWidth(w interface{}) (fixedWidth, bool) {
	a, ok := w.(fixedWidth)
	if !ok || a.getWidth() <= 0 {
		return nil, false
	}
	return a, ok
}

type fixedHeight interface {
	getHeight() float64
}

func isFixedHeight(w interface{}) (fixedHeight, bool) {
	a, ok := w.(fixedHeight)
	if !ok || a.getHeight() <= 0 {
		return nil, false
	}
	return a, ok
}
