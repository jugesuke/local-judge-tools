package scol

import (
	"fmt"
)

type Color struct {
	R uint8
	G uint8
	B uint8
}

var (
	WHITE = Color{R: 255, G: 255, B: 255}
	BLACK = Color{R: 0, G: 0, B: 0}
)

func SetColor(fg, bg Color, s string) string {
	str := fmt.Sprintf("\x1b[38;2;%d;%d;%d;48;2;%d;%d;%d;1m%s\x1b[m", fg.R, fg.G, fg.B, bg.R, bg.G, bg.B, s)
	return str
}
