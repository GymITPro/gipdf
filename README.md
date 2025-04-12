# gipdf

A flexible PDF generation library for Go, built on top of `github.com/phpdave11/gofpdf`. This library provides a component-based API for creating PDF documents with a fluid builder pattern.

## Installation

```bash
go get github.com/GymITPro/gipdf
```

## Features

- Component-based document building
- Headers and footers
- Text fields with custom formatting
- Data fields with labels
- Image embedding
- Signature fields
- Custom fonts
- Automatic page breaks
- Layout management with rows and columns

## Basic Usage

```go
package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/GymITPro/gipdf"
)

func main() {
	// Create a new document with configuration
	doc := gipdf.NewDocument(gipdf.Config{
		Padding: gipdf.PaddingLTRB(10, 10, 10, 10),
		Fonts: []*gipdf.ConfigFont{
			{
				Name:  "Helvetica",
				Style: "",
				Data:  nil, // Built-in font
			},
		},
	})

	// Add a header
	doc.AddHeader(gipdf.PaddingAll(5), 2, func(r *gipdf.Row) {
		r.TextField("Document Header", 1, 10, gipdf.AlignmentCenter, 
			gipdf.UseFontSize(14))
	})

	// Add content
	doc.Row(gipdf.PaddingAll(5), 2, func(r *gipdf.Row) {
		r.TextField("Hello World", 1, 10, gipdf.AlignmentLeft)
	})

	// Generate the PDF
	pdfBytes, err := doc.Render()
	if err != nil {
		log.Fatal(err)
	}

	// Save to file
	err = ioutil.WriteFile("output.pdf", pdfBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Components

### Document

The main container for your PDF content.

```go
doc := gipdf.NewDocument(gipdf.Config{
    Padding: gipdf.PaddingAll(10),
})
```

### Rows and Columns

Create layouts using rows (horizontal) and columns (vertical).

```go
doc.Row(gipdf.PaddingAll(5), 2, func(r *gipdf.Row) {
    r.Column(gipdf.PaddingAll(2), 1, 0, 1, func(c *gipdf.Column) {
        // Column content
    })
})
```

### Widgets

Add content to your PDF with various widgets:

#### TextField

```go
row.TextField("This is text", 1, 10, gipdf.AlignmentLeft, 
    gipdf.UseFontSize(12),
    gipdf.UseColor(gipdf.ColorRGB(0, 0, 0)))
```

#### DataField

```go
row.DataField("Label", "Value", 1, 8, 10)
```

#### Image

```go
imageData, _ := ioutil.ReadFile("logo.png")
row.Image(imageData, "png", 1)
```

#### SignatureField

```go
signatureData, _ := ioutil.ReadFile("signature.png")
row.SignatureField(signatureData, "png", "2022-01-01", "John Doe", 1, 10)
```

#### EmptyField

```go
row.EmptyField(1)
```

### Headers and Footers

```go
// Add a standard header
doc.AddHeader(gipdf.PaddingAll(5), 2, func(r *gipdf.Row) {
    r.TextField("Document Header", 1, 10, gipdf.AlignmentCenter)
})

// Add a first page header (different from other pages)
doc.AddFirstPageHeader(gipdf.PaddingAll(5), 2, func(r *gipdf.Row) {
    r.TextField("First Page Header", 1, 10, gipdf.AlignmentCenter)
})

// Add a footer
doc.AddFooter(gipdf.PaddingAll(5), 2, func(r *gipdf.Row) {
    r.TextField("Page " + gipdf.CurrentPage() + " of " + gipdf.TotalPages(), 
        1, 10, gipdf.AlignmentCenter)
})
```

### Styling

Apply styling to text using configuration functions:

```go
row.TextField("Styled Text", 1, 10, gipdf.AlignmentLeft,
    gipdf.UseFont(gipdf.Font{Name: "Helvetica", Style: "B", Size: 14}),
    gipdf.UseColor(gipdf.ColorRGB(255, 0, 0)),
    gipdf.UseFillColor(gipdf.ColorRGB(240, 240, 240)))
```

## Page Breaks

Add manual page breaks:

```go
doc.PageBreak()
```

## License

MIT