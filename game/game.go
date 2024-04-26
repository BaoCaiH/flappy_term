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

const (
	start = 0 + iota
	active
	over
)

type Game struct {
	screen      *screen.Screen
	bird        *bird.Bird
	plumber     []*plumber.Pipe
	inputEvents chan int
	score       int
	ticker      time.Ticker
	state       int
}

func Init() *Game {
	return &Game{
		screen:      screen.Init(Width, Height),
		bird:        bird.Init(),
		plumber:     plumber.Init(),
		inputEvents: make(chan int),
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
	g.screen.Render(left, top, fg, bg)

	if g.state == start {
		utils.Draw(left+1, top+2, "##### #       #   ####  ####  #   #", termbox.ColorYellow, bg)
		utils.Draw(left+1, top+3, "#     #      # #  #   # #   #  # #", termbox.ColorYellow, bg)
		utils.Draw(left+1, top+4, "##### #     ##### ####  ####    #", termbox.ColorYellow, bg)
		utils.Draw(left+1, top+5, "#     #     #   # #     #       #", termbox.ColorYellow, bg)
		utils.Draw(left+1, top+6, "#     ##### #   # #     #       #", termbox.ColorYellow, bg)

		utils.Draw(left+7, top+8, "##### ##### ####  #   #", termbox.ColorLightRed, bg)
		utils.Draw(left+7, top+9, "  #   #     #   # ## ##", termbox.ColorLightRed, bg)
		utils.Draw(left+7, top+10, "  #   ##### ####  # # #", termbox.ColorLightRed, bg)
		utils.Draw(left+7, top+11, "  #   #     #  #  #   #", termbox.ColorLightRed, bg)
		utils.Draw(left+7, top+12, "  #   ##### #   # #   #", termbox.ColorLightRed, bg)

		utils.Draw(left+1, top+18, "Press [Space] or [k] or [Up] to Jump", termbox.ColorYellow, bg)
		utils.Draw(left+4, top+20, "Press [Enter] or [i] to Start", termbox.ColorGreen, bg)
		utils.Draw(left+5, top+22, "Press [Esc] or [q] to Quit", termbox.ColorRed, bg)

		return termbox.Flush()
	}

	if g.state == over {
		g.screen.Render(left, top, fg, bg)
		utils.Draw(left+7, top+2, " ####   #   #   # #####", termbox.ColorLightRed, bg)
		utils.Draw(left+7, top+3, "#      # #  ## ## #", termbox.ColorLightRed, bg)
		utils.Draw(left+7, top+4, "# ### ##### # # # #####", termbox.ColorLightRed, bg)
		utils.Draw(left+7, top+5, "#   # #   # #   # #", termbox.ColorLightRed, bg)
		utils.Draw(left+7, top+6, " ###  #   # #   # #####", termbox.ColorLightRed, bg)

		utils.Draw(left+7, top+8, " ###  #   # ##### ####", termbox.ColorLightRed, bg)
		utils.Draw(left+7, top+9, "#   # #   # #     #   #", termbox.ColorLightRed, bg)
		utils.Draw(left+7, top+10, "#   # #   # ##### ####", termbox.ColorLightRed, bg)
		utils.Draw(left+7, top+11, "#   #  # #  #     #  #", termbox.ColorLightRed, bg)
		utils.Draw(left+7, top+12, " ###    #   ##### #   #", termbox.ColorLightRed, bg)

		scoreText := fmt.Sprintf("SCORE: %d", g.score)
		utils.Draw(left+(Width-len(scoreText))/2, top+16, scoreText, fg, bg)

		utils.Draw(left+1, top+18, "Press [Space] or [k] or [Up] to Jump", termbox.ColorYellow, bg)
		utils.Draw(left+4, top+20, "Press [Enter] or [i] to Start", termbox.ColorGreen, bg)
		utils.Draw(left+5, top+22, "Press [Esc] or [q] to Quit", termbox.ColorRed, bg)

		return termbox.Flush()
	}

	g.bird.Render(left, top, bg)
	for _, p := range g.plumber {
		p.Render(top, bottom, left, right, bg)
	}
	utils.Draw(right-21, bottom+1, "[Esc] or [q] to Quit", fg, bg)
	utils.Draw(left, bottom+1, fmt.Sprintf("Score: %d", g.score), fg, bg)

	return termbox.Flush()
}

// Tick in an expected interval to handle physics
func (g *Game) physicsLoop(stopped chan bool) {
	for {
		if g.state != active {
			continue
		}
		select {
		case <-stopped:
			return
		case <-g.ticker.C:
			// Move pipes
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

// Game loop will tick when ever it can
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
			case input.Start:
				g.score = 0
				g.bird.Reset()
				plumber.Reset(g.plumber)
				g.state = active
			case input.Jump:
				g.bird.Moving = -3
			case input.End:
				g.state = over
			default:
			}
		default:
		}

		if err = g.render(); err != nil {
			panic(err)
		}
	}
}
