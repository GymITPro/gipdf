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

	document.Row(PaddingAll(4), 0.75, func(row *Row) {
		row.TextFieldFixed("Test", 20, 0, 10, AlignmentLeft, UseBackgroundColor(ColorRGB(0, 255, 0)))
		row.TextField("Test 1", 1, 10, AlignmentLeft, UseBackgroundColor(ColorRGB(255, 0, 0)))
		row.TextField("Test 2", 1, 10, AlignmentLeft, UseBackgroundColor(ColorRGB(0, 0, 255)))
	})

	document.Row(PaddingAll(8), 4, func(row *Row) {
		row.Column(PaddingAll(0), 4, 20, 1, func(column *Column) {
			column.TextFieldFixed("Test 2",
				10,
				10,
				10,
				AlignmentLeft,
				UseFont(Font{
					Name:  "Arial",
					Style: "B",
					Size:  24,
				}),
				UseBackgroundColor(ColorRGB(0, 0, 0)),
				UseColor(ColorRGB(255, 255, 255)),
			)
			column.TextField("Test 3",
				1,
				10,
				AlignmentLeft,
				UseFont(Font{
					Name:  "Arial",
					Style: "",
					Size:  14,
				}),
				UseColor(ColorRGB(158, 170, 193)))
		})
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
