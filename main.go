package main

import (
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
	defer screen.Fini()

	game := NewGame()

	ticker := time.NewTicker(800 * time.Millisecond)
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
		drawGame(screen, game)
		screen.Show()
	}
}

func drawGame(screen tcell.Screen, g *game) {
	for y, row := range g.board {
		for x, color := range row {
			screen.SetContent(x*2, y, ' ', []rune{'.'}, tcell.StyleDefault.Background(color.tcellColor()).Foreground(color.tcellColor()))
		}
	}

	drawPiece(screen, g.offset.x, g.offset.y, g.curPiece)
}

func drawPiece(screen tcell.Screen, x, y int, t *tetromino) {
	for _, coord := range t.Coords() {
		screen.SetContent((x+coord.dx)*2, y+coord.dy, ' ', []rune{' '}, tcell.StyleDefault.Background(t.color.tcellColor()))
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
