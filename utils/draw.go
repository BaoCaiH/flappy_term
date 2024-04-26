package utils

import "github.com/nsf/termbox-go"

func Draw(x, y int, msg string, fg, bg termbox.Attribute) {
	for _, ch := range msg {
		termbox.SetCell(x, y, ch, fg, bg)
		x += 1
	}
}
