package main

import "github.com/gdamore/tcell"

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
		screen.SetContent(x+coord.dx*2, y+coord.dy, ' ', nil, tcell.StyleDefault.Background(t.color.tcellColor()))
		screen.SetContent(x+coord.dx*2+1, y+coord.dy, ' ', nil, tcell.StyleDefault.Background(t.color.tcellColor()))
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
