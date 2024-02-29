package gipdf

import (
	"bytes"
)

type signatureField struct {
	Image       []byte       `json:"image"`
	Date        string       `json:"date"`
	Subtitle    string       `json:"subtitle"`
	ImageType   string       `json:"image_type"`
	Font        *Font        `json:"font"`
	TextHeight  float64      `json:"text_height"`
	AspectRatio float64      `json:"aspect_ratio"`
	Configs     []configFunc `json:"-"`
}

func (r *Row) SignatureField(image []byte, imageType, date, subtitle string, aspectRatio, textHeight float64, configs ...configFunc) *Row {
	r.Columns = append(r.Columns, signatureField{
		Image:       image,
		Date:        date,
		Subtitle:    subtitle,
		ImageType:   imageType,
		TextHeight:  textHeight,
		Configs:     configs,
		AspectRatio: aspectRatio,
	})
	return r
}

func (r *Column) SignatureField(image []byte, imageType, date, subtitle string, aspectRatio, textHeight float64, configs ...configFunc) *Column {
	r.Rows = append(r.Rows, signatureField{
		Image:       image,
		Date:        date,
		Subtitle:    subtitle,
		ImageType:   imageType,
		TextHeight:  textHeight,
		Configs:     configs,
		AspectRatio: aspectRatio,
	})
	return r
}

func (c signatureField) Render(pdf *Document, x, y, width, height float64) {
	imageUrl := randomString()
	pdf.RegisterImageReader(imageUrl, c.ImageType, bytes.NewReader(c.Image))
	pdf.Image(imageUrl, x, y, width, 0, false, "", 0, "")
	info := pdf.GetImageInfo(imageUrl)
	h := (width / info.Width()) * info.Height()
	pdf.Ln(h)
	for _, config := range c.Configs {
		config(pdf)
	}

	pdf.CellFormat(width, c.TextHeight, c.Date, "", 1, "", false, 0, "")
	pdf.Line(x, pdf.GetY(), x+width, pdf.GetY())
	pdf.CellFormat(width, c.TextHeight, c.Subtitle, "", 1, "", false, 0, "")
}

func (c signatureField) GetAspectRatio() float64 {
	return c.AspectRatio
}
