package gipdf

type Page struct {
	Elements   []Element
	Margin     Margin
	BackGround *Color
	Header     Element
	Footer     Element
}

func (p *Page) HeaderHeight() float64 {
	if p.Header != nil {
		if h := p.Header.FixedHeight(); h != nil {
			return *h
		}
	}
	return 0
}

func (p *Page) FooterHeight() float64 {
	if p.Footer != nil {
		if h := p.Footer.FixedHeight(); h != nil {
			return *h
		}
	}
	return 0
}

func (p *Page) AspectRatio() float64  { return 1 }
func (p *Page) FixedWidth() *float64  { return nil }
func (p *Page) FixedHeight() *float64 { return nil }
func (p *Page) Render(ctx *RenderContext, x, y, width, height float64) error {
	for _, el := range p.Elements {
		err := el.Render(ctx, x, ctx.CursorY, width, height)
		if err != nil {
			return err
		}
	}
	return nil
}
