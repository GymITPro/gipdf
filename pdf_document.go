package gipdf

import (
	"bytes"

	"github.com/phpdave11/gofpdf"
)

const maxWidth = 210

type Document struct {
	Padding         Padding `json:"padding"`
	Rows            []*Row  `json:"rows"`
	FirstPageHeader []*Row  `json:"first_page_header"`
	Header          []*Row  `json:"header"`
	Footer          []*Row  `json:"footer"`
	Spacing         float64 `json:"spacing"`
	width           float64
	defaultFont     Font
	*gofpdf.Fpdf
}

type Header struct {
	Title    string `json:"title"`
	SubTitle string `json:"sub_title"`
	Image    []byte `json:"image"`
}

func NewDocument(config Config) *Document {
	pdf := gofpdf.New("P", "mm", "A4", "")
	var defaultFont *Font
	for _, font := range config.Fonts {
		pdf.AddUTF8FontFromBytes(font.Name, font.Style, font.Data)
		if font.Default != nil {
			pdf.SetFont(font.Name, font.Style, *font.Default)
			defaultFont = &Font{
				Name:  font.Name,
				Style: font.Style,
				Size:  *font.Default,
			}
		}
	}

	if defaultFont == nil {
		defaultFont = &Font{
			Name:  "helvetica",
			Style: "",
			Size:  12,
		}

		pdf.SetFont(defaultFont.Name, defaultFont.Style, defaultFont.Size)
	}

	pdf.SetMargins(config.Padding.Left, config.Padding.Top, config.Padding.Right)
	pdf.SetAutoPageBreak(true, config.Padding.Bottom)
	return &Document{
		Padding:         config.Padding,
		Spacing:         config.Spacing,
		Rows:            nil,
		FirstPageHeader: nil,
		Header:          nil,
		Footer:          nil,
		defaultFont:     *defaultFont,
		width:           maxWidth - config.Padding.Left - config.Padding.Right,
		Fpdf:            pdf,
	}
}

func GetHeight(row *Row) float64 {
	if a, ok := isFixedHeight(row); ok {
		return a.getHeight()
	}

	document := NewDocument(Config{
		Padding: PaddingAll(0),
	})
	document.Ln(0)
	row.render(document, 0, 0, document.width, 0)
	return document.GetY()
}

func (d *Document) AddPage() {
	d.Fpdf.AddPage()
}

func (d *Document) Render() ([]byte, error) {
	d.header()
	d.footer()
	d.AddPage()
	d.header()
	for _, row := range d.Rows {
		// Auto page break
		height := GetHeight(row)
		_, pageHeight := d.GetPageSize()
		if pageHeight-d.Padding.Bottom < d.GetY()+height {
			d.AddPage()
		}

		row.render(d, d.GetX(), d.GetY(), d.width, height)
		if d.Fpdf.Err() {
			return nil, d.Fpdf.Error()
		}
		if d.GetY()+d.Spacing > pageHeight-d.Padding.Bottom {
			d.AddPage()
		} else {
			d.Ln(d.Spacing)
		}
	}

	var buffer bytes.Buffer
	err := d.Output(&buffer)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (d *Document) header() {
	if d.FirstPageHeader != nil && d.PageNo() == 0 {
		d.SetHeaderFuncMode(func() {
			d.SetX(0)
			d.SetY(0)
			var maxY float64 = 0
			for _, row := range d.FirstPageHeader {
				d.Ln(0)
				row.render(d, d.GetX(), d.GetY(), d.width, 0)
				if d.GetY() > maxY {
					maxY = d.GetY()
				}
			}
			d.SetY(maxY + d.Padding.Top)
			d.SetX(d.Padding.Left)
		}, false)
		return
	}

	d.SetHeaderFuncMode(func() {
		d.SetX(0)
		d.SetY(0)
		var maxY float64 = 0
		for _, row := range d.Header {
			d.Ln(0)
			row.render(d, d.GetX(), d.GetY(), d.width, 0)
			if d.GetY() > maxY {
				maxY = d.GetY()
			}
		}
		d.SetY(maxY + d.Padding.Top)
		d.SetX(d.Padding.Left)
	}, false)
}

func (d *Document) AddHeader(padding Padding, spacing float64, builder func(*Row)) *Document {

	header := &Row{
		Columns:     nil,
		Padding:     padding,
		Spacing:     spacing,
		IsPageBreak: false,
		AspectRatio: 0,
		builder:     builder,
	}
	d.Header = append(d.Header, header)
	return d
}

func (d *Document) AddFirstPageHeader(padding Padding, spacing float64, builder func(*Row)) *Document {

	header := &Row{
		Columns:     nil,
		Padding:     padding,
		Spacing:     spacing,
		IsPageBreak: false,
		AspectRatio: 0,
		builder:     builder,
	}
	d.FirstPageHeader = append(d.FirstPageHeader, header)
	return d
}

func (d *Document) AddFooter(padding Padding, spacing float64, builder func(*Row)) *Document {

	footer := &Row{
		Columns:     nil,
		Padding:     padding,
		Spacing:     spacing,
		IsPageBreak: false,
		AspectRatio: 0,
		builder:     builder,
	}
	d.Footer = append(d.Footer, footer)
	return d
}

func (d *Document) footer() {
	d.SetFooterFunc(func() {
		d.SetY(-15)
		for _, row := range d.Footer {
			d.Ln(0)
			row.render(d, d.GetX(), d.GetY(), d.width, 0)
		}
	})
	d.AliasNbPages("")
}
