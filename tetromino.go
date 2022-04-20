package main

type Dir int

const (
	Up Dir = iota
	Right
	Down
	Left
)

type Color int

const (
	Cyan Color = iota
	Blue
	Orange
	Yellow
	Green
	Purple
	Red
)

type Tetromino struct {
	color  Color
	blocks [4]uint16
}

var TETROMINOS = [7]Tetromino{
	{color: Cyan, blocks: [4]uint16{0x0F00, 0x2222, 0x00F0, 0x4444}},
	{color: Blue, blocks: [4]uint16{0x44C0, 0x8E00, 0x6440, 0x0E20}},
	{color: Orange, blocks: [4]uint16{0x4460, 0x0E80, 0xC440, 0x2E00}},
	{color: Yellow, blocks: [4]uint16{0xCC00, 0xCC00, 0xCC00, 0xCC00}},
	{color: Green, blocks: [4]uint16{0x06C0, 0x8C40, 0x6C00, 0x4620}},
	{color: Purple, blocks: [4]uint16{0x0E40, 0x4C40, 0x4E00, 0x4640}},
	{color: Red, blocks: [4]uint16{0x0C60, 0x4C80, 0xC600, 0x2640}},
}
