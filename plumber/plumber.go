package plumber

import (
	"flappy-term/utils"
	"math/rand/v2"
	"strings"

	"github.com/nsf/termbox-go"
)

type Pipe struct {
	X, Y int
}

const (
	pipeLength = 20
	pipeGap    = 7
	pipeWidth  = 5
	pipeDist   = 11
)

//	func Init(startX, startY int) *Pipe {
//		return &Pipe{X: startX, Y: startY}
//	}
func Init() []*Pipe {
	return []*Pipe{
		{42, rand.IntN(12) - 12},
		{42 + pipeWidth + pipeDist, rand.IntN(12) - 12},
		{42 + (pipeWidth+pipeDist)*2, rand.IntN(12) - 12},
	}
}

func (p *Pipe) Render(top, bottom, left, right int, bg termbox.Attribute) {
	pipeLeft := p.X + left
	pipeRight := p.X + left + pipeWidth
	showableWidth := pipeWidth
	if pipeLeft < left {
		showableWidth -= min(left-pipeLeft, showableWidth)
	} else if pipeRight > right-1 {
		showableWidth -= min(pipeRight-right+1, showableWidth)
	}
	for i := 0; i < 20; i++ {
		if i == 19 {
			utils.Draw(max(pipeLeft, left), p.Y+i, strings.Repeat("#", showableWidth), termbox.ColorGreen, bg)
		} else if i == 0 {
			utils.Draw(max(pipeLeft, left), p.Y+pipeLength+pipeGap, strings.Repeat("#", showableWidth), termbox.ColorGreen, bg)
		} else {
			if p.Y+i > top {
				if pipeLeft >= left && pipeLeft < right-1 {
					termbox.SetCell(pipeLeft, p.Y+i, '#', termbox.ColorGreen, bg)
				}
				if pipeLeft+pipeWidth-1 >= left && pipeLeft+pipeWidth-1 < right-1 {
					termbox.SetCell(pipeLeft+pipeWidth-1, p.Y+i, '#', termbox.ColorGreen, bg)
				}
			}
			if p.Y+i+pipeLength+pipeGap < bottom {
				if pipeLeft >= left && pipeLeft < right-1 {
					termbox.SetCell(pipeLeft, p.Y+i+pipeLength+pipeGap, '#', termbox.ColorGreen, bg)
				}
				if pipeLeft+pipeWidth-1 >= left && pipeLeft+pipeWidth-1 < right-1 {
					termbox.SetCell(pipeLeft+pipeWidth-1, p.Y+i+pipeLength+pipeGap, '#', termbox.ColorGreen, bg)
				}
			}
		}
	}
}

func Reset(pipes []*Pipe) {
	for i, p := range pipes {
		p.X = 42 + i*(pipeWidth+pipeDist)
	}
}

func (p *Pipe) Reset() {
	p.X = pipeWidth*2 + pipeDist*3
	p.Y = rand.IntN(12) - 12
}
