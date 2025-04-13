package gipdf

import "github.com/signintech/gopdf"

type Unit int

const (
	UnitUnset Unit = gopdf.UnitUnset // No units were set, when conversion is called on nothing will happen
	UnitPT    Unit = gopdf.UnitPT    // Points
	UnitMM    Unit = gopdf.UnitMM    // Millimeters
	UnitCM    Unit = gopdf.UnitCM    // Centimeters
	UnitIN    Unit = gopdf.UnitIN    // Inches
	UnitPX    Unit = gopdf.UnitPX    // Pixels
)

var PageSizeA4 = gopdf.PageSizeA4

type Config struct {
	PageSize        *gopdf.Rect
	PageUnits       Unit
	Fonts           []*Font
	BackGroundColor *Color
	Debug           bool
}
