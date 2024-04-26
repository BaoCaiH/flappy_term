package arena

import (
	"flappy-term/utils"

	"github.com/nsf/termbox-go"
)

type Arena struct {
	W, H int
	Pos  utils.Vec2
}

func Init(w, h int) *Arena {
	return &Arena{
		W: w,
		H: h,
	}
}

func (a *Arena) Render(x, y int, fg, bg termbox.Attribute) {
	top := y
	bottom := y + a.H + 1
	left := x
	for i := top; i < bottom; i++ {
		termbox.SetCell(left-1, i, '|', fg, bg)
		termbox.SetCell(left+a.W, i, '|', fg, bg)
	}

	termbox.SetCell(left-1, top, '┌', fg, bg)
	termbox.SetCell(left-1, bottom, '└', fg, bg)
	termbox.SetCell(left+a.W, top, '┐', fg, bg)
	termbox.SetCell(left+a.W, bottom, '┘', fg, bg)

	a.fill(left, top, a.W, 1, '─', fg, bg)
	a.fill(left, bottom, a.W, 1, '─', fg, bg)

}

func (a *Arena) fill(x, y, w, h int, ch rune, fg, bg termbox.Attribute) {
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			termbox.SetCell(x+col, y+row, ch, fg, bg)
		}
	}
}
