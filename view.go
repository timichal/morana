/*
methods handling the view = termbox frontend
*/

package main

import (
	"github.com/nsf/termbox-go"
	"strconv"
)

func initView() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

func (view *View) stop() {
	termbox.Close()
}

func (view *View) refresh() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// drawing the main screen
	view.drawMap("0")
	// drawing the bottom bar
	view.drawBar()

	view.drawDebug(debugtext)

	termbox.Flush()

	// input handling
	keyInput.handleInput()
}

func (view *View) drawMap(floor string) {
	// draw map by tile type
	for x, row := range floormap[floor] {
		for y, tile := range row {
			termbox.SetCell(x, y, tile.TileType, termbox.ColorGreen, termbox.ColorBlack)
		}
	}

	// draw the player
	termbox.SetCell(player.xpos, player.ypos, '@', termbox.ColorRed, termbox.ColorBlack)
}

func (view *View) drawBar() {
	bartext := "Name: " + player.Name + " | Level " + strconv.Itoa(player.Level) + " | HP " + strconv.Itoa(player.HP) + " | Moves " + strconv.Itoa(moves)
	view.drawText(0, 20, bartext, termbox.ColorWhite, termbox.ColorBlack)
}

func (view *View) drawText(x int, y int, text string, fg termbox.Attribute, bg termbox.Attribute) {
	for index, char := range text {
		termbox.SetCell(x+index, y, rune(char), fg, bg)
	}
}

// debug line
func (view *View) drawDebug(text string) {
	for index, char := range text {
		termbox.SetCell(0+index, 21, rune(char), termbox.ColorRed, termbox.ColorBlack)
	}
}
