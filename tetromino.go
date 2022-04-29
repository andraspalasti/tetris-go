package main

import (
	"math/rand"

	"github.com/gdamore/tcell"
)

type Color int

const (
	None Color = iota
	Cyan
	Blue
	Orange
	Yellow
	Green
	Purple
	Red
)

type tetromino struct {
	color  Color
	blocks [4]uint16
	rot    int
}

var TETROMINOS = [7]tetromino{
	{color: Cyan, blocks: [4]uint16{0x0F00, 0x2222, 0x00F0, 0x4444}},
	{color: Blue, blocks: [4]uint16{0x44C0, 0x8E00, 0x6440, 0x0E20}},
	{color: Orange, blocks: [4]uint16{0x4460, 0x0E80, 0xC440, 0x2E00}},
	{color: Yellow, blocks: [4]uint16{0xCC00, 0xCC00, 0xCC00, 0xCC00}},
	{color: Green, blocks: [4]uint16{0x06C0, 0x8C40, 0x6C00, 0x4620}},
	{color: Purple, blocks: [4]uint16{0x0E40, 0x4C40, 0x4E00, 0x4640}},
	{color: Red, blocks: [4]uint16{0x0C60, 0x4C80, 0xC600, 0x2640}},
}

// Returns a new random tetromino
func NewTetromino() *tetromino {
	piece := TETROMINOS[rand.Intn(len(TETROMINOS))]
	return &piece
}

// Returns the blocks of the tetromino represented as a uint16
func (t *tetromino) Blocks() uint16 {
	return t.blocks[t.rot]
}

// Rotates the tetromino to the right
func (t *tetromino) Rotate() {
	t.rot = (t.rot + 1) % len(t.blocks)
}

// Rotates the tetromino to the left
func (t *tetromino) RotateBack() {
	t.rot--
	if t.rot < 0 {
		t.rot = len(t.blocks) - 1
	}
}

// Returns the coordinates of the tetrominos blocks
func (t *tetromino) Coords() []struct{ dx, dy int } {
	blocks := t.Blocks()
	dx, dy := 0, 0

	var coords []struct{ dx, dy int }
	for i := uint16(0x8000); i != 0; i = i >> 1 {
		if (blocks & i) != 0 {
			coords = append(coords, struct {
				dx int
				dy int
			}{dx, dy})
		}
		if dx++; dx == 4 {
			dx = 0
			dy++
		}
	}

	return coords
}

// Returns the color as a tcell.Color
func (c *Color) tcellColor() tcell.Color {
	switch *c {
	case Cyan:
		return tcell.ColorDarkCyan
	case Blue:
		return tcell.ColorBlue
	case Orange:
		return tcell.ColorOrange
	case Yellow:
		return tcell.ColorYellow
	case Green:
		return tcell.ColorGreen
	case Purple:
		return tcell.ColorDarkMagenta
	case Red:
		return tcell.ColorRed
	default:
		return tcell.ColorDefault
	}
}
