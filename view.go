/*
methods handling the view = termbox frontend
*/

package main

import (
	"github.com/nsf/termbox-go"
	"os"
	"strconv"
)

func initView() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	view.refresh()
}

func (view *View) stop() {
	termbox.Close()
}

func (view *View) refresh() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// input handling
	switch engine.State {
	case "Intro":
		view.drawIntro()
	case "GameOn":
		// drawing the main screen
		view.drawMap(engine.CurrentFloor)
		// drawing the bottom bar
		view.drawBar()

		view.drawDebug(debugtext)
	case "Victory":
		view.drawVictoryLine()
	default:
		view.stop()
		os.Exit(0)
	}
	termbox.Flush()
}

func (view *View) drawIntro() {
	introtext_1 :=`___  ___                            `
	introtext_2 :=`|  \/  |                            `
	introtext_3 :=`| .  . | ___  _ __ __ _ _ __   __ _ `
	introtext_4 :=`| |\/| |/ _ \| '__/ _| | '_ \ / _| |`
	introtext_5 :=`| |  | | (_) | | | (_| | | | | (_| |`
	introtext_6 :=`\_|  |_/\___/|_|  \__,_|_| |_|\__,_|`
	view.drawText(10, 10, introtext_1, termbox.ColorWhite, termbox.ColorBlack)
	view.drawText(10, 11, introtext_2, termbox.ColorWhite, termbox.ColorBlack)
	view.drawText(10, 12, introtext_3, termbox.ColorWhite, termbox.ColorBlack)
	view.drawText(10, 13, introtext_4, termbox.ColorWhite, termbox.ColorBlack)
	view.drawText(10, 14, introtext_5, termbox.ColorWhite, termbox.ColorBlack)
	view.drawText(10, 15, introtext_6, termbox.ColorWhite, termbox.ColorBlack)
}

func (view *View) drawVictoryLine() {
	bartext := "You win! Press r to restart or esc/q to quit"
	view.drawText(0, 20, bartext, termbox.ColorWhite, termbox.ColorBlack)
}

func (view *View) drawMap(floor string) {
	// draw map by tile type
	for x, row := range floormap[floor] {
		for y, tile := range row {
			termbox.SetCell(x, y, tile.TileType, termbox.ColorGreen, termbox.ColorBlack)
		}
	}

	// draw the player
	termbox.SetCell(player.PosX, player.PosY, '@', termbox.ColorRed, termbox.ColorBlack)
}

func (view *View) drawBar() {
	bartext := "Name: " + player.Name + " | Level " + strconv.Itoa(player.Level) + " | HP " + strconv.Itoa(player.HP) + " | Moves " + strconv.Itoa(player.Moves)
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
