package gipdf

import (
	"io"

	"github.com/pkg/errors"
	"github.com/signintech/gopdf"
)

func New(config *Config) *PDF {
	return &PDF{
		pdf:     &gopdf.GoPdf{},
		PageSet: &PageSet{},
		Config:  config,
	}
}

func (p *PDF) AddPage(elements ...Element) *Page {
	page := &Page{
		Elements:   elements,
		BackGround: p.Config.BackGroundColor,
	}

	p.PageSet.Pages = append(p.PageSet.Pages, page)
	return page
}
func (p *PDF) Render() error {
	pageSize := *PageSizeA4
	if p.Config.PageSize != nil {
		pageSize = *p.Config.PageSize
	}

	p.pdf.Start(gopdf.Config{PageSize: pageSize})

	for _, font := range p.Config.Fonts {
		data, err := font.Data()
		if err != nil {
			return errors.Wrapf(err, "failed to load font data for %s", font.Family)
		}
		if err := p.pdf.AddTTFFontData(font.Family, data); err != nil {
			return errors.Wrapf(err, "failed to add font %s", font.Family)
		}

		if font.Default {
			if err := p.pdf.AddTTFFontData("", data); err != nil {
				return errors.Wrapf(err, "failed to add font as default %s", font.Family)
			}
		}
	}

	return p.PageSet.Render(p.pdf, pageSize, p.Config.Debug)
}

func (p *PDF) WriteTo(w io.Writer) (int64, error) {
	return p.pdf.WriteTo(w)
}

func (p *PDF) WriteFile(filename string) error {
	return p.pdf.WritePdf(filename)
}
