package bird

import (
	"flappy-term/utils"

	"github.com/nsf/termbox-go"
)

type Bird struct {
	Pos    utils.Vec2
	Hitbox utils.Vec2
}

func Init(startX, startY int) *Bird {
	return &Bird{Pos: utils.Vec2{X: startX, Y: startY}, Hitbox: utils.Vec2{X: 3, Y: 2}}
}

func (b *Bird) Render(x, y int, bg termbox.Attribute) {
	top := y + b.Pos.Y
	left := x + b.Pos.X

	// First line
	utils.Draw(left, top, "◢■", termbox.ColorLightRed, bg)
	termbox.SetCell(left+2, top, 'Ꙩ', termbox.ColorWhite, bg)
	// Second line
	termbox.SetCell(left, top+1, '◀', termbox.ColorWhite, bg)
	utils.Draw(left+1, top+1, "■■", termbox.ColorLightRed, bg)
	termbox.SetCell(left+3, top+1, '=', termbox.ColorYellow, bg)
}
