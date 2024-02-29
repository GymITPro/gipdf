package gipdf

type Config struct {
	Padding Padding       `json:"padding"`
	Fonts   []*ConfigFont `json:"fonts"`
}

type ConfigFont struct {
	Name    string   `json:"name"`
	Style   string   `json:"style"`
	Data    []byte   `json:"data"`
	Default *float64 `json:"default"`
}

type Font struct {
	Name  string  `json:"name"`
	Style string  `json:"style"`
	Size  float64 `json:"size"`
}

type configFunc func(*Document)

type Color struct {
	Red   int `json:"red"`
	Green int `json:"green"`
	Blue  int `json:"blue"`
}

func ColorRGB(r, g, b int) Color {
	return Color{
		Red:   r,
		Green: g,
		Blue:  b,
	}
}
