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
		drawBox(screen, 0, 0, game.Width()*2+2, game.Height()+2)
		drawBox(screen, int(float32(game.Width()*2)*1.5), 0, 10, 5)
		drawGame(screen, 1, 1, game)
		screen.Show()

		if game.hasLost {

		}
	}
}

func drawGame(screen tcell.Screen, x, y int, g *game) {
	for i, row := range g.board {
		for j, color := range row {
			screen.SetContent(x+j*2, y+i, ' ', []rune{'.'}, tcell.StyleDefault.Background(color.tcellColor()).Foreground(color.tcellColor()))
		}
	}

	drawPiece(screen, x+g.offset.x*2, y+g.offset.y, g.curPiece)
}

func drawPiece(screen tcell.Screen, x, y int, t *tetromino) {
	for _, coord := range t.Coords() {
		screen.SetContent(x+coord.dx*2, y+coord.dy, ' ', []rune{' '}, tcell.StyleDefault.Background(t.color.tcellColor()))
	}
}

func drawBox(screen tcell.Screen, x, y int, w, h int) {
	if w < 0 || h < 0 {
		panic("Width or height of box cannot be less than 0")
	}

	// draw corners
	screen.SetContent(x, y, tcell.RuneULCorner, nil, tcell.StyleDefault)
	screen.SetContent(x+w-1, y, tcell.RuneURCorner, nil, tcell.StyleDefault)
	screen.SetContent(x, y+h-1, tcell.RuneLLCorner, nil, tcell.StyleDefault)
	screen.SetContent(x+w-1, y+h-1, tcell.RuneLRCorner, nil, tcell.StyleDefault)

	// draw the two rows
	for i := 1; i < w-1; i++ {
		screen.SetContent(x+i, y, tcell.RuneHLine, nil, tcell.StyleDefault)
		screen.SetContent(x+i, y+h-1, tcell.RuneHLine, nil, tcell.StyleDefault)
	}

	// draw the two columns
	for i := 1; i < h-1; i++ {
		screen.SetContent(x, y+i, tcell.RuneVLine, nil, tcell.StyleDefault)
		screen.SetContent(x+w-1, y+i, tcell.RuneVLine, nil, tcell.StyleDefault)
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
