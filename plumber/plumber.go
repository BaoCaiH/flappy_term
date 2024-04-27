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
	PipeGap   = 7
	PipeWidth = 5
	pipeDist  = 11
)

func Init() []*Pipe {
	return []*Pipe{
		{42, rand.IntN(24-6-PipeGap) + 3},
		{42 + PipeWidth + pipeDist, rand.IntN(24-6-PipeGap) + 3},
		{42 + (PipeWidth+pipeDist)*2, rand.IntN(24-6-PipeGap) + 3},
	}
}

func (p *Pipe) Render(top, bottom, left, right int, bg termbox.Attribute) {
	pipeLeft := p.X + left
	pipeRight := p.X + left + PipeWidth
	showableWidth := PipeWidth
	if pipeLeft < left {
		showableWidth -= min(left-pipeLeft, showableWidth)
	} else if pipeRight > right-1 {
		showableWidth -= min(pipeRight-right+1, showableWidth)
	}
	for i := top + 1; i < bottom; i++ {
		if i == p.Y+top {
			utils.Draw(max(pipeLeft, left), i, strings.Repeat("#", showableWidth), termbox.ColorGreen, bg)
		} else if i == p.Y+top+PipeGap+1 {
			utils.Draw(max(pipeLeft, left), i, strings.Repeat("#", showableWidth), termbox.ColorGreen, bg)
		} else if i < p.Y+top || i > p.Y+top+PipeGap+1 {
			if pipeLeft >= left && pipeLeft < right-1 {
				termbox.SetCell(pipeLeft, i, '#', termbox.ColorGreen, bg)
			}
			if pipeLeft+PipeWidth-1 >= left && pipeLeft+PipeWidth-1 < right-1 {
				termbox.SetCell(pipeLeft+PipeWidth-1, i, '#', termbox.ColorGreen, bg)
			}
		}
	}
}

func Reset(pipes []*Pipe) {
	for i, p := range pipes {
		p.X = 42 + i*(PipeWidth+pipeDist)
	}
}

func (p *Pipe) Reset() {
	p.X = PipeWidth*2 + pipeDist*3
	p.Y = rand.IntN(24-6-PipeGap) + 3
}
