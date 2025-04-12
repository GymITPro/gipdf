package gipdf

import (
	"bytes"
)

type columnImage struct {
	Image       []byte       `json:"image"`
	ImageType   string       `json:"image_type"`
	AspectRatio float64      `json:"aspect_ratio"`
	FixedHeight float64      `json:"fixed_height"`
	FixedWidth  float64      `json:"fixed_width"`
	configs     []ConfigFunc `json:"-"`
}

func (r *Row) Image(image []byte, imageType string, aspectRatio float64, configs ...ConfigFunc) *Row {
	r.Columns = append(r.Columns, columnImage{
		Image:       image,
		AspectRatio: aspectRatio,
		ImageType:   imageType,
		configs:     configs,
	})
	return r
}

func (r *Row) ImageFixed(image []byte, imageType string, width, height float64, configs ...ConfigFunc) *Row {
	r.Columns = append(r.Columns, columnImage{
		Image:       image,
		FixedWidth:  width,
		FixedHeight: height,
		ImageType:   imageType,
		configs:     configs,
	})
	return r
}

func (r *Column) Image(image []byte, imageType string, aspectRatio float64, configs ...ConfigFunc) *Column {
	r.Rows = append(r.Rows, columnImage{
		Image:       image,
		AspectRatio: aspectRatio,
		ImageType:   imageType,
		configs:     configs,
	})
	return r
}

func (r *Column) ImageFixed(image []byte, imageType string, width, height float64, configs ...ConfigFunc) *Column {
	r.Rows = append(r.Rows, columnImage{
		Image:       image,
		FixedWidth:  width,
		FixedHeight: height,
		ImageType:   imageType,
		configs:     configs,
	})
	return r
}

func (c columnImage) render(pdf *Document, x, y, width, height float64) {
	imageUrl := randomString()
	pdf.RegisterImageReader(imageUrl, c.ImageType, bytes.NewReader(c.Image))
	height = 0.0
	if c.FixedHeight > 0 {
		height = c.FixedHeight
	}
	pdf.Image(imageUrl, x, y, width, height, false, "", 0, "")
	info := pdf.GetImageInfo(imageUrl)
	h := (width / info.Width()) * info.Height()
	pdf.Ln(h)

}

func (c columnImage) getAspectRatio() float64 {
	return c.AspectRatio
}

func (c columnImage) getHeight() float64 {
	return c.FixedHeight
}

func (c columnImage) getWidth() float64 {
	return c.FixedWidth
}
