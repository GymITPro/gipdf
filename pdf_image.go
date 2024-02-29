package gipdf

import (
	"bytes"
)

type columnImage struct {
	Image       []byte  `json:"image"`
	ImageType   string  `json:"image_type"`
	AspectRatio float64 `json:"aspect_ratio"`
}

func (r *Row) Image(image []byte, imageType string, aspectRatio float64) *Row {
	r.Columns = append(r.Columns, columnImage{
		Image:       image,
		AspectRatio: aspectRatio,
		ImageType:   imageType,
	})
	return r
}

func (r *Column) Image(image []byte, imageType string, aspectRatio float64) *Column {
	r.Rows = append(r.Rows, columnImage{
		Image:       image,
		AspectRatio: aspectRatio,
		ImageType:   imageType,
	})
	return r
}

func (c columnImage) Render(pdf *Document, x, y, width, height float64) {
	imageUrl := randomString()
	pdf.RegisterImageReader(imageUrl, c.ImageType, bytes.NewReader(c.Image))
	pdf.Image(imageUrl, x, y, width, 0, false, "", 0, "")
	info := pdf.GetImageInfo(imageUrl)
	h := (width / info.Width()) * info.Height()
	pdf.Ln(h)

}

func (c columnImage) GetAspectRatio() float64 {
	return c.AspectRatio
}
