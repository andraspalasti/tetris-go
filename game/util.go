package game

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

var colorToAttr = map[Color]termbox.Attribute{
	Cyan:   termbox.ColorCyan,
	Blue:   termbox.ColorBlue,
	Orange: termbox.ColorYellow,
	Yellow: termbox.ColorLightYellow,
	Green:  termbox.ColorGreen,
	Purple: termbox.ColorMagenta,
	Red:    termbox.ColorRed,
	None:   termbox.ColorDefault,
}
