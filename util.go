package gipdf

import (
	"math"

	"github.com/signintech/gopdf"
)

func Rectangle(gp *gopdf.GoPdf, x0 float64, y0 float64, x1 float64, y1 float64, style string, radi [4]float64, radiusPointNum int) error {
	if x1 <= x0 || y1 <= y0 {
		return gopdf.ErrInvalidRectangleCoordinates
	}
	allEmpty := true
	for _, radius := range radi {
		if radius != 0 {
			allEmpty = false
		}
	}

	if radiusPointNum <= 0 || allEmpty {
		// draw rectangle without round corner
		var points []gopdf.Point
		points = append(points, gopdf.Point{X: x0, Y: y0})
		points = append(points, gopdf.Point{X: x1, Y: y0})
		points = append(points, gopdf.Point{X: x1, Y: y1})
		points = append(points, gopdf.Point{X: x0, Y: y1})
		gp.Polygon(points, style)

	} else {
		for _, radius := range radi {
			if radius > (x1-x0) || radius > (y1-y0) {
				return gopdf.ErrInvalidRectangleCoordinates
			}
		}

		var degrees []float64
		angle := float64(90) / float64(radiusPointNum+1)
		accAngle := angle
		for accAngle < float64(90) {
			degrees = append(degrees, accAngle)
			accAngle += angle
		}

		var radians []float64
		for _, v := range degrees {
			radians = append(radians, v*math.Pi/180)
		}

		var points = []gopdf.Point{}
		points = append(points, gopdf.Point{X: x0, Y: y0 + radi[0]})
		for _, v := range radians {
			offsetX := radi[0] * math.Cos(v)
			offsetY := radi[0] * math.Sin(v)
			x := x0 + radi[0] - offsetX
			y := y0 + radi[0] - offsetY
			points = append(points, gopdf.Point{X: x, Y: y})
		}
		points = append(points, gopdf.Point{X: x0 + radi[0], Y: y0})

		points = append(points, gopdf.Point{X: x1 - radi[1], Y: y0})
		for i := range radians {
			v := radians[len(radians)-1-i]
			offsetX := radi[1] * math.Cos(v)
			offsetY := radi[1] * math.Sin(v)
			x := x1 - radi[1] + offsetX
			y := y0 + radi[1] - offsetY
			points = append(points, gopdf.Point{X: x, Y: y})
		}
		points = append(points, gopdf.Point{X: x1, Y: y0 + radi[1]})

		points = append(points, gopdf.Point{X: x1, Y: y1 - radi[2]})
		for _, v := range radians {
			offsetX := radi[2] * math.Cos(v)
			offsetY := radi[2] * math.Sin(v)
			x := x1 - radi[2] + offsetX
			y := y1 - radi[2] + offsetY
			points = append(points, gopdf.Point{X: x, Y: y})
		}
		points = append(points, gopdf.Point{X: x1 - radi[2], Y: y1})

		points = append(points, gopdf.Point{X: x0 + radi[3], Y: y1})
		for i := range radians {
			v := radians[len(radians)-1-i]
			offsetX := radi[3] * math.Cos(v)
			offsetY := radi[3] * math.Sin(v)
			x := x0 + radi[3] - offsetX
			y := y1 - radi[3] + offsetY
			points = append(points, gopdf.Point{X: x, Y: y})
		}
		points = append(points, gopdf.Point{X: x0, Y: y1 - radi[3]})

		gp.Polygon(points, style)
	}
	return nil
}
