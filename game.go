package main

type game struct {
	score    uint
	curPiece *tetromino
	offset   struct{ x, y int }
	board    [][]Color
}

func NewGame() *game {
	width, height := 10, 20

	g := game{
		score:    0,
		curPiece: NewTetromino(),
		board:    make([][]Color, height),
		offset:   struct{ x, y int }{width/2 -1, 0}}

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

// Moves the current piece in the game to the left
// The move only occurs if there is space for the piece
func (g *game) MoveLeft() {
	if !g.isOccupied(g.offset.x-1, g.offset.y) {
		g.offset.x--
	}
}

// Moves the current piece in the game to the right
// The move only occurs if there is space for the piece
func (g *game) MoveRight() {
	if !g.isOccupied(g.offset.x+1, g.offset.y) {
		g.offset.x++
	}
}

// Moves the current piece in the game down
// The move only occurs if there is space for the piece
func (g *game) MoveDown() {
	if !g.isOccupied(g.offset.x, g.offset.y+1) {
		g.offset.y++
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

	// move blocks from tetromino to the board
	for _, coord := range g.curPiece.Coords() {
		g.board[g.offset.y+coord.dy][g.offset.x+coord.dx] = g.curPiece.color
	}

	g.removeFullRows()
	g.curPiece = NewTetromino()
	g.offset.x, g.offset.y = len(g.board[0])/2-1, 0
}

func (g *game) removeFullRows() {
rowLoop:
	for i, row := range g.board {
		// if its not a full row continue
		for _, block := range row {
			if block == None {
				continue rowLoop
			}
		}

		// shift all upward rows downward
		for j := i - 1; 0 <= j; j-- {
			g.board[j+1] = g.board[j]
		}
		g.board[0] = make([]Color, len(g.board[0]))
	}
}

// If there is space for the current piece it returns true else it returns false.
// The x and y parameters are for specifing the offset's of the piece
func (g *game) isOccupied(x, y int) bool {
	w, h := len(g.board[0]), len(g.board)
	for _, coord := range g.curPiece.Coords() {
		cur_x, cur_y := x+coord.dx, y+coord.dy

		if cur_y < 0 || cur_x < 0 || h <= cur_y || w <= cur_x || g.board[cur_y][cur_x] != None {
			return true
		}
	}
	return false
}
