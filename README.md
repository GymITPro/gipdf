# gipdf

A Go library for creating PDF documents with an element-based composition system. Built on top of [github.com/signintech/gopdf](https://github.com/signintech/gopdf) with an easier API for common document creation tasks.

## Features

- Composable element-based architecture
- Simple layout management with support for dimensions and aspect ratios
- Box elements with background colors and rounded corners
- Image elements with automatic sizing
- Debug mode to visualize element boundaries
- Easy extension for custom elements

## Installation

```bash
go get github.com/yourusername/gipdf
```

## Basic Usage

```go
package main

import (
    "github.com/yourusername/gipdf"
)

func main() {
    // Create a new document
    doc := gipdf.New(gipdf.Config{
        Width:  595.28,
        Height: 841.89,
        Debug:  false, // Set to true to visualize element boundaries
    })

    // Add a page
    doc.AddPage()

    // Create a box with a blue background
    box := &gipdf.Box{
        Width:      gipdf.Float(400),
        Height:     gipdf.Float(200),
        Background: gipdf.Color{R: 66, G: 133, B: 244},
        CornerRadius: [4]float64{10, 10, 10, 10}, // Rounded corners
    }

    // Add the box to the document at position x:50, y:50
    doc.Add(box, 50, 50)

    // Add an image
    img := &gipdf.Image{
        Path:  "path/to/image.jpg",
        Width: gipdf.Float(200), // Fixed width
        Ratio: 16.0/9.0,         // Will determine height based on aspect ratio
    }

    // Add the image inside a box
    boxWithImage := &gipdf.Box{
        Width:      gipdf.Float(300),
        Height:     gipdf.Float(200),
        Background: gipdf.Color{R: 255, G: 255, B: 255},
        Children:   []gipdf.Element{img},
    }

    doc.Add(boxWithImage, 50, 300)

    // Save the document
    doc.Save("output.pdf")
}
```

## Element Types

### Box

A container element that can have a background color, rounded corners, and contain child elements.

```go
box := &gipdf.Box{
    Width:        gipdf.Float(300),
    Height:       gipdf.Float(200),
    Background:   gipdf.Color{R: 240, G: 240, B: 240},
    CornerRadius: [4]float64{5, 5, 5, 5}, // TL, TR, BR, BL
    Children:     []gipdf.Element{/* child elements */},
}
```

### Image

An element for rendering images with control over dimensions and aspect ratio.

```go
image := &gipdf.Image{
    Path:   "path/to/image.jpg",
    Ratio:  1.5,               // Aspect ratio (width/height)
    Width:  gipdf.Float(200),  // Fixed width (optional)
    Height: gipdf.Float(150),  // Fixed height (optional)
}
```

## Advanced Usage

### Debug Mode

Enable debug mode to see element boundaries:

```go
doc := gipdf.New(gipdf.Config{
    Width:  595.28,
    Height: 841.89,
    Debug:  true,
})
```

### Creating Custom Elements

Implement the Element interface to create custom elements:

```go
type Element interface {
    Render(ctx *RenderContext, x, y, width, height float64) error
    AspectRatio() float64
    FixedWidth() *float64
    FixedHeight() *float64
}
```

Example custom element:

```go
type Circle struct {
    Diameter *float64
    Color    Color
}

func (c *Circle) AspectRatio() float64  { return 1.0 }
func (c *Circle) FixedWidth() *float64  { return c.Diameter }
func (c *Circle) FixedHeight() *float64 { return c.Diameter }

func (c *Circle) Render(ctx *RenderContext, x, y, width, height float64) error {
    // Implementation to draw a circle
    // ...
    return nil
}
```

## License

MIT

## Dependencies

- [github.com/signintech/gopdf](https://github.com/signintech/gopdf) - Base PDF generation library

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.