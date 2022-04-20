package main

type Game struct {
	score    uint
	curPiece *Tetromino
	offset   struct{ x, y uint }
	board    [][]Color
}

func NewGame(width, height uint) *Game {
	if width < 10 || height < 10 {
		panic("Width and height of the board can not be less than 10")
	}

	game := Game{
		score:    0,
		curPiece: NewTetromino(),
		board:    make([][]Color, height),
		offset:   struct{ x, y uint }{width/2 - 2, 0}}

	for i := range game.board {
		game.board[i] = make([]Color, width)
	}
	return &game
}

func (g *Game) MovePiece(dir Dir) {
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

func (g *Game) RotatePiece() {
	g.curPiece.Rotate()
	if g.isOccupied(g.offset.x, g.offset.y) {
		g.curPiece.RotateBack()
	}
}

func (g *Game) RemoveFullRows() {
	for i, row := range g.board {
		// if its not a full row
		// it will jump to the next item in the loop
		for _, block := range row {
			if block == None {
				continue
			}
		}

		// shift all upward rows downward
		for j := i - 1; 0 <= j; j-- {
			g.board[j+1] = g.board[j]
		}
		g.board[0] = make([]Color, len(g.board[0]))
	}
}

func (g *Game) isOccupied(x uint, y uint) bool {
	blocks := g.curPiece.Blocks()
	dx, dy := uint(0), uint(0)
	w, h := uint(len(g.board[0])), uint(len(g.board))

	for i := uint16(0xFFFF); i != 0; i = i >> 1 {
		if (blocks & i) != 0 {
			if w <= y+dy || h <= x+dx || g.board[y+dy][x+dx] != None {
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
