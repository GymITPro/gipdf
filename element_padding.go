package gipdf

type Padding struct {
	Top, Right, Bottom, Left float64
}

type PaddingElement struct {
	Element Element
	Padding Padding
}

func (p *PaddingElement) Render(ctx *RenderContext, x, y, width, height float64) error {
	innerX := x + p.Padding.Left
	innerY := y + p.Padding.Top
	innerW := width - p.Padding.Left - p.Padding.Right
	innerH := height - p.Padding.Top - p.Padding.Bottom

	return p.Element.Render(ctx, innerX, innerY, innerW, innerH)
}

func (p *PaddingElement) FixedWidth() *float64 {
	if fw := p.Element.FixedWidth(); fw != nil {
		w := *fw + p.Padding.Left + p.Padding.Right
		return &w
	}
	return nil
}

func (p *PaddingElement) FixedHeight() *float64 {
	if fh := p.Element.FixedHeight(); fh != nil {
		h := *fh + p.Padding.Top + p.Padding.Bottom
		return &h
	}
	return nil
}

func (p *PaddingElement) AspectRatio() float64 {
	return p.Element.AspectRatio()
}

func Pad(element Element, top, right, bottom, left float64) *PaddingElement {
	return &PaddingElement{
		Element: element,
		Padding: Padding{Top: top, Right: right, Bottom: bottom, Left: left},
	}
}
