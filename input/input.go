package input

import "github.com/nsf/termbox-go"

const (
	Quit int = -1 + iota
	Start
	Jump
)

func HandleInput(inputChan chan int) {
	for {
		switch event := termbox.PollEvent(); {
		case event.Type == termbox.EventError:
			panic(event.Err)
		case event.Key == termbox.KeyEsc || event.Key == termbox.KeyCtrlC || event.Ch == 'q':
			inputChan <- Quit
		case event.Key == termbox.KeyEnter || event.Ch == 'i':
			inputChan <- Start
		case event.Key == termbox.KeySpace || event.Ch == 'k' || event.Key == termbox.KeyArrowUp:
			inputChan <- Jump
		}
	}
}
