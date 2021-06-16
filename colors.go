package mdtopdf

import "strconv"

var (
	Black = Color{0, 0, 0}
	White = Color{255, 255, 255}
)

// Grey returns a grey of some level between 0 and 255.
func Grey(level int) Color {
	return Color{
		Red:   level,
		Green: level,
		Blue:  level,
	}
}

// ColorOf implements basic CSS-like colours such as "#0366d6".
func ColorOf(css string) Color {
	if css == "" {
		return Black
	}

	if css[0] == '#' {
		if len(css) == 4 {
			r, _ := strconv.ParseInt(css[1:2], 16, 64)
			g, _ := strconv.ParseInt(css[2:3], 16, 64)
			b, _ := strconv.ParseInt(css[3:], 16, 64)
			return Color{Red: int(r), Green: int(g), Blue: int(b)}
		}
		if len(css) == 7 {
			r, _ := strconv.ParseInt(css[1:3], 16, 64)
			g, _ := strconv.ParseInt(css[3:5], 16, 64)
			b, _ := strconv.ParseInt(css[5:], 16, 64)
			return Color{Red: int(r), Green: int(g), Blue: int(b)}
		}
	}

	return Black
}
