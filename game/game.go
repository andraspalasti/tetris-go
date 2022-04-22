package game

import "github.com/nsf/termbox-go"

type game struct {
	score    uint
	curPiece *Tetromino
	offset   struct{ x, y int }
	board    [][]Color
}

func NewGame() *game {
	width, height := 10, 20

	g := game{
		score:    0,
		curPiece: NewTetromino(),
		board:    make([][]Color, height),
		offset:   struct{ x, y int }{width/2 - 2, 0}}

	for i := range g.board {
		g.board[i] = make([]Color, width)
	}

	return &g
}

func (g *game) Score() uint {
	return g.score
}

func (g *game) Height() int {
	return len(g.board)
}

func (g *game) Width() int {
	return len(g.board[0])
}

func (g *game) MovePiece(dir Dir) {
	switch dir {
	case Left:
		if !g.isOccupied(g.offset.x-1, g.offset.y) {
			g.offset.x--
		}
	case Right:
		if !g.isOccupied(g.offset.x+1, g.offset.y) {
			g.offset.x++
		}
	case Down:
		if !g.isOccupied(g.offset.x, g.offset.y+1) {
			g.offset.y++
		}
	}
}

func (g *game) Draw() {
	for y, row := range g.board {
		for x, color := range row {
			termbox.SetCell(x*2, y, ' ', colorToAttr[color], colorToAttr[color])
			termbox.SetCell(x*2+1, y, '.', colorToAttr[color], colorToAttr[color])
		}
	}

	blocks := g.curPiece.Blocks()
	dx, dy := 0, 0
	for i := uint16(0x8000); i != 0; i = i >> 1 {
		if (blocks & i) != 0 {
			x, y := g.offset.x+dx, g.offset.y+dy
			termbox.SetCell(x*2, y, ' ', colorToAttr[g.curPiece.color], colorToAttr[g.curPiece.color])
			termbox.SetCell(x*2+1, y, ' ', colorToAttr[g.curPiece.color], colorToAttr[g.curPiece.color])
		}
		if dx++; dx == 4 {
			dx = 0
			dy++
		}
	}
}

func (g *game) RotatePiece() {
	g.curPiece.Rotate()
	if g.isOccupied(g.offset.x, g.offset.y) {
		g.curPiece.RotateBack()
	}
}

func (g *game) Tick() {
	// if we can move the piece down
	if !g.isOccupied(g.offset.x, g.offset.y+1) {
		g.offset.y++
		return
	}

	// move blocks from tetromino from the board
	blocks := g.curPiece.Blocks()
	dx, dy := 0, 0
	for i := uint16(0x8000); i != 0; i = i >> 1 {
		if (blocks & i) != 0 {
			g.board[g.offset.y+dy][g.offset.x+dx] = g.curPiece.color
		}
		if dx++; dx == 4 {
			dx = 0
			dy++
		}
	}

	g.removeFullRows()
	g.curPiece = NewTetromino()
	g.offset.x, g.offset.y = len(g.board[0])/2-2, 0
}

func (g *game) removeFullRows() {
	for i, row := range g.board {
		isFull := true
		for _, block := range row {
			if block == None {
				isFull = false
				break
			}
		}

		if isFull {
			// shift all upward rows downward
			for j := i - 1; 0 <= j; j-- {
				g.board[j+1] = g.board[j]
			}
			g.board[0] = make([]Color, len(g.board[0]))
		}
	}
}

func (g *game) isOccupied(x, y int) bool {
	blocks := g.curPiece.Blocks()
	dx, dy := 0, 0
	w, h := len(g.board[0]), len(g.board)

	for i := uint16(0x8000); i != 0; i = i >> 1 {
		if (blocks & i) != 0 {
			cur_x, cur_y := x+dx, y+dy
			if cur_y < 0 || cur_x < 0 || h <= cur_y || w <= cur_x || g.board[cur_y][cur_x] != None {
				return true
			}
		}
		if dx++; dx == 4 {
			dx = 0
			dy++
		}
	}
	return false
}
