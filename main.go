package main

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	if err := screen.Init(); err != nil {
		panic(err)
	}
	screen.DisableMouse()
	defer screen.Fini()

	rand.Seed(time.Now().UnixNano())
	game := NewGame()

	ticker := time.NewTicker(600 * time.Millisecond)
	events := listenEvents(screen)

gameLoop:
	for {
		select {
		case <-ticker.C:
			game.Tick()
		case ev := <-events:
			switch ev.Key() {
			case tcell.KeyEsc:
				break gameLoop
			case tcell.KeyUp:
				game.RotatePiece()
			case tcell.KeyRight:
				game.MoveRight()
			case tcell.KeyLeft:
				game.MoveLeft()
			case tcell.KeyDown:
				game.MoveDown()
			}
		}
		screen.Clear()
		screen.Show()
		drawBox(screen, 0, 0, game.Width()*2+2, game.Height()+2)
		drawBox(screen, int(float32(game.Width()*2)*1.5), 0, 12, 6)
		drawPiece(screen, int(float32(game.Width()*2)*1.5)+1, 1, game.nextPiece)
		drawGame(screen, 1, 1, game)

		if game.hasLost {
		}

		screen.Show()
	}
}

func listenEvents(screen tcell.Screen) <-chan *tcell.EventKey {
	events := make(chan *tcell.EventKey)
	go func() {
		for {
			switch ev := screen.PollEvent().(type) {
			case *tcell.EventResize:
				screen.Sync()
			case *tcell.EventKey:
				events <- ev
			}
		}
	}()
	return events
}
