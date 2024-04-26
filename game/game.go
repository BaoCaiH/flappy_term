package game

import (
	"flappy-term/bird"
	"flappy-term/input"
	"flappy-term/plumber"
	"flappy-term/screen"
	"flappy-term/utils"
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	Width  = 37 // supposed to be 16, only results in 2/3 height
	Height = 24
	fg     = termbox.ColorDefault
	bg     = termbox.ColorDefault
)

type Game struct {
	screen      *screen.Screen
	bird        *bird.Bird
	plumber     []*plumber.Pipe
	inputEvents chan int
	score       int
	ticker      time.Ticker
	// pipe        *plumber.Pipe
}

func Init() *Game {
	return &Game{
		screen:      screen.Init(Width, Height),
		bird:        bird.Init(8, 12),
		plumber:     plumber.Init(),
		inputEvents: make(chan int),
		// pipe:        plumber.Init(36, -12),
	}
}

func (g *Game) render() error {
	termbox.Clear(fg, bg)

	w, h := termbox.Size()
	midHeight := h / 2
	top := midHeight - g.screen.H/2
	bottom := top + g.screen.H + 1
	left := (w - g.screen.W) / 2
	right := left + g.screen.W + 1

	// Render name and score because too much work to have a separate class, oops
	utils.Draw(left, top-1, "Flappy Term", fg, bg)
	utils.Draw(left, bottom+1, fmt.Sprintf("Score: %d", g.score), fg, bg)

	g.screen.Render(left, top, fg, bg)
	g.bird.Render(left, top, bg)
	for _, p := range g.plumber {
		p.Render(top, bottom, left, right, bg)
	}

	return termbox.Flush()
}

func (g *Game) physicsLoop(stopped chan bool) {
	for {
		select {
		case <-stopped:
			return
		case <-g.ticker.C:
			// Move pipe
			for _, p := range g.plumber {
				if p.X == -5 {
					p.Reset()
				} else {
					p.X -= 1
				}
				if p.X == g.bird.X {
					g.score += 1
				}
			}
			// Move bird
			g.bird.Y = min(max(g.bird.Y+g.bird.Moving, 0), g.screen.H)
			g.bird.Moving = 1
		}
	}
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

	if duration, err := time.ParseDuration("125ms"); err != nil {
		panic(err)
	} else {
		g.ticker = *time.NewTicker(duration)
	}

	tickerStopped := make(chan bool)
	go g.physicsLoop(tickerStopped)

gameLoop:
	for {
		select {
		case event := <-g.inputEvents:
			switch event {
			case input.Quit:
				g.ticker.Stop()
				break gameLoop
				// case input.Start:
			case input.Jump:
				g.bird.Moving = -3
			default:
			}
		default:
		}

		if err = g.render(); err != nil {
			panic(err)
		}
	}
}
