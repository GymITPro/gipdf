package gipdf

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	fontData, err := os.ReadFile("testdata/FiraSans-Regular.ttf")
	if err != nil {
		t.Fatal(err)
	}

	pdf := New(&Config{
		PageSize: PageSizeA4,
		Fonts: []*Font{
			{
				Family:  "FiraSans",
				Data:    func() ([]byte, error) { return fontData, nil },
				Default: true,
			},
		},
		BackGroundColor: &Color{
			R: 244,
			G: 246,
			B: 247,
		},
		Debug: true,
	})

	pdf.AddPage(&Column{
		Elements: []Element{
			&Row{
				Ratio: 1,
				Elements: []Element{
					&Box{
						Ratio: 1,
						Background: Color{
							R: 158,
							G: 170,
							B: 193,
						},
						CornerRadius: [4]float64{
							0,  // TL
							20, // TR
							20, // BR
							0,  // BL
						},
						Children: []Element{
							&Text{
								Text:     "Hello World",
								FontSize: 12,
							},
						},
					},
					&Empty{
						Ratio: 1,
					},
				},
			},
			&Text{
				Height:   Ptr(20.0),
				Text:     "Hello World",
				FontSize: 12,
			},
		},
	})

	err = pdf.Render()
	if err != nil {
		t.Fatal(err)
	}

	err = pdf.WriteFile("testdata/test.pdf")
	if err != nil {
		t.Fatal(err)
	}

}
