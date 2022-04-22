package main

import (
	"github.com/andraspalasti/tetris/game"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"time"
)

func main() {
	// init termbox to control terminal
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// go routine for listening to events
	inputEvents := make(chan termbox.Event)
	go func() {
		for {
			ev := termbox.PollEvent()
			inputEvents <- ev
		}
	}()

	g := game.NewGame()

	// game loop
	start := time.Now()
	for {
		dTime := time.Now().Sub(start)

		var ev termbox.Event
		select {
		case ev = <-inputEvents:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				g.MovePiece(game.Left)

			case termbox.KeyArrowRight:
				g.MovePiece(game.Right)

			case termbox.KeyArrowDown:
				g.MovePiece(game.Down)

			case termbox.KeyArrowUp:
				g.RotatePiece()
			}
		default: // no event
		}

		if ev.Key == termbox.KeyEsc || ev.Ch == 'q' {
			break
		}

		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		g.Draw()
		termbox.Flush()

		// make it tick here
		if 1000 < dTime.Milliseconds() {
			g.Tick()
			start = time.Now()
		}
	}
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}
