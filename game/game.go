package game

import (
	"flappy-term/arena"
	"flappy-term/bird"
	"flappy-term/input"
	"flappy-term/utils"
	"fmt"

	"github.com/nsf/termbox-go"
)

const (
	Width  = 37 // supposed to be 16, only results in 2/3 height
	Height = 24
	fg     = termbox.ColorDefault
	bg     = termbox.ColorDefault
)

type Game struct {
	arena       *arena.Arena
	bird        *bird.Bird
	inputEvents chan int
	score       int
}

func Init() *Game {
	return &Game{arena: arena.Init(Width, Height), bird: bird.Init(8, 12), inputEvents: make(chan int)}
}

func (g *Game) render() error {
	termbox.Clear(fg, bg)

	w, h := termbox.Size()
	midHeight := h / 2
	top := midHeight - g.arena.H/2
	bottom := midHeight + g.arena.H/2 + 1
	left := (w - g.arena.W) / 2

	// Render name and score because too much work to have a separate class, oops
	utils.Draw(left, top-1, "Flappy Term", fg, bg)
	utils.Draw(left, bottom+1, fmt.Sprintf("Score: %d", g.score), fg, bg)

	g.arena.Render(left, top, fg, bg)
	g.bird.Render(left, top, bg)

	return termbox.Flush()
}

func (g *Game) Start() {
	var err error
	if err = termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	go input.HandleInput(g.inputEvents)

	if err = g.render(); err != nil {
		panic(err)
	}

gameLoop:
	for {
		select {
		case event := <-g.inputEvents:
			switch event {
			case input.Quit:
				break gameLoop
				// case input.Start:
			}
		}

		if err = g.render(); err != nil {
			panic(err)
		}
	}
}
