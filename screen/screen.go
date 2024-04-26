package screen

import (
	"github.com/nsf/termbox-go"
)

type Screen struct {
	W, H int
}

func Init(w, h int) *Screen {
	return &Screen{
		W: w,
		H: h,
	}
}

func (s *Screen) Render(x, y int, fg, bg termbox.Attribute) {
	top := y
	bottom := y + s.H + 1
	left := x
	for i := top; i < bottom; i++ {
		termbox.SetCell(left-1, i, '|', fg, bg)
		termbox.SetCell(left+s.W, i, '|', fg, bg)
	}

	termbox.SetCell(left-1, top, '┌', fg, bg)
	termbox.SetCell(left-1, bottom, '└', fg, bg)
	termbox.SetCell(left+s.W, top, '┐', fg, bg)
	termbox.SetCell(left+s.W, bottom, '┘', fg, bg)

	s.fill(left, top, s.W, 1, '─', fg, bg)
	s.fill(left, bottom, s.W, 1, '─', fg, bg)

}

func (s *Screen) fill(x, y, w, h int, ch rune, fg, bg termbox.Attribute) {
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			termbox.SetCell(x+col, y+row, ch, fg, bg)
		}
	}
}
