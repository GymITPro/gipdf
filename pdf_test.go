package gipdf

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	document := NewDocument(Config{
		Padding: PaddingAll(0),
		Spacing: 0,
	})

	document.Row(PaddingAll(0), 4, func(row *Row) {
		row.Column(PaddingAll(10), 4, 50, 1, func(column *Column) {
			column.TextField("Test 2",
				1,
				24,
				AlignmentLeft,
				UseFont(Font{
					Name:  "Arial",
					Style: "B",
					Size:  24,
				}),
				UseColor(ColorRGB(0, 0, 0)),
			)
			column.TextField("Test 3",
				1,
				14,
				AlignmentLeft,
				UseFont(Font{
					Name:  "Arial",
					Style: "",
					Size:  14,
				}),
				UseColor(ColorRGB(158, 170, 193)))
		})
		row.TextFieldFixed("Test 1", 20, 20, 10, AlignmentRight)
	})

	render, err := document.Render()
	if err != nil {
		t.Fatal(err)
	}

	err = os.MkdirAll("testdata", 0755)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile("testdata/test.pdf", render, 0644)
	if err != nil {
		t.Fatal(err)
	}
}
