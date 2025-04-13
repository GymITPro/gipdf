package gipdf

// Empty is a placeholder for empty elements in the document.
type Empty struct {
	Ratio  float64
	Width  *float64
	Height *float64
}

func (e *Empty) AspectRatio() float64  { return e.Ratio }
func (e *Empty) FixedWidth() *float64  { return e.Width }
func (e *Empty) FixedHeight() *float64 { return e.Height }
func (e *Empty) Render(ctx *RenderContext, x, y, width, height float64) error {
	// No rendering needed for empty elements
	return nil
}
