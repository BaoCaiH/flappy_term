package bird

import (
	"flappy-term/utils"

	"github.com/nsf/termbox-go"
)

type Bird struct {
	X, Y, HX, HY int
	// Magnitude of "gravity" relative to the screen, 1 means move down 1 line
	Moving int
}

func Init() *Bird {
	b := Bird{}
	b.Reset()
	return &b
}

func (b *Bird) Render(x, y int, bg termbox.Attribute) {
	top := y + b.Y
	left := x + b.X

	// First line
	utils.Draw(left, top, "◢■", termbox.ColorYellow, bg)
	termbox.SetCell(left+2, top, 'Ꙩ', termbox.ColorWhite, bg)
	// Second line
	termbox.SetCell(left, top+1, '◀', termbox.ColorWhite, bg)
	utils.Draw(left+1, top+1, "■■", termbox.ColorYellow, bg)
	termbox.SetCell(left+3, top+1, '=', termbox.ColorLightRed, bg)
}

func (b *Bird) Reset() {
	b.X = 8
	b.Y = 12
}
