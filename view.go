/*
methods handling the view = termbox frontend
*/

package main

import (
	"github.com/nsf/termbox-go"
	"os"
	"strconv"
	"unicode/utf8"
)

func initView() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	view.width = viewWidth
	view.height = viewHeight

	view.refresh()
}

func (view *View) stop() {
	termbox.Close()
}

func (view *View) refresh() {
	debugtext := "This is the debug bar"
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// input handling
	switch engine.State {
	case "Intro":
		view.drawIntro()
	case "GameOn":
		view.drawTopBar()
		// drawing the main screen
		view.drawMap(engine.CurrentFloor)
		// drawing the bottom bar
		view.drawBottomBar()

		view.drawDebugBar(debugtext)
	case "Victory":
		view.drawVictoryLine()
	default:
		view.stop()
		os.Exit(0)
	}
	termbox.Flush()
}

func (view *View) drawIntro() {
	for i := 0; i < view.width; i++ {
		view.drawText(i, 0, " ", termbox.ColorBlack, termbox.ColorYellow)
		view.drawText(i, view.height-1, " ", termbox.ColorBlack, termbox.ColorYellow)
	}
	for i := 0; i < view.height; i++ {
		view.drawText(0, i, " ", termbox.ColorBlack, termbox.ColorYellow)
		view.drawText(view.width-1, i, " ", termbox.ColorBlack, termbox.ColorYellow)
	}

	introLogo := [6]string{
		"███╗   ███╗ ██████╗ ██████╗  █████╗ ███╗   ██╗ █████╗ ",
		"████╗ ████║██╔═══██╗██╔══██╗██╔══██╗████╗  ██║██╔══██╗",
		"██╔████╔██║██║   ██║██████╔╝███████║██╔██╗ ██║███████║",
		"██║╚██╔╝██║██║   ██║██╔══██╗██╔══██║██║╚██╗██║██╔══██║",
		"██║ ╚═╝ ██║╚██████╔╝██║  ██║██║  ██║██║ ╚████║██║  ██║",
		"╚═╝     ╚═╝ ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝"}

	logoHeight := view.height/2 - (len(introLogo) / 2) - 2
	logoWidth := view.width/2 - utf8.RuneCountInString(introLogo[0])/2
	for i := 0; i < 6; i++ {
		view.drawText(logoWidth, logoHeight+i, introLogo[i], termbox.ColorWhite, termbox.ColorBlack)
	}

	descText := "A roguelite with a nice intro screen and not much else"
	view.drawText(logoWidth, logoHeight+len(introLogo)+2, descText, termbox.ColorWhite, termbox.ColorBlack)

	instText := "Press any key to start, q to quit"
	view.drawText(logoWidth+10, logoHeight+len(introLogo)+4, instText, termbox.ColorWhite, termbox.ColorBlack)
}

func (view *View) drawVictoryLine() {
	bartext := "You win! Press r to restart or esc/q to quit"
	view.drawText(0, 20, bartext, termbox.ColorWhite, termbox.ColorBlack)
}

func (view *View) drawTopBar() {
	bartext := "MORANA"
	view.drawText(0, 0, bartext, termbox.ColorWhite, termbox.ColorBlack)
}

func (view *View) drawMap(floor string) {
	// draw map by tile type
	for x, row := range floormap[floor] {
		for y, tile := range row {
			termbox.SetCell(x, y+1, tile.TileType, termbox.ColorGreen, termbox.ColorBlack)
		}
	}

	// draw the player
	termbox.SetCell(player.PosX, player.PosY+1, '@', termbox.ColorRed, termbox.ColorBlack)
}

func (view *View) drawBottomBar() {
	bartext := "Name: " + player.Name + " | Level " + strconv.Itoa(player.Level) + " | HP " + strconv.Itoa(player.HP) + " | Moves " + strconv.Itoa(player.Moves)
	bartext2 := "I guess something will be here as well"
	view.drawText(0, view.height-3, bartext, termbox.ColorWhite, termbox.ColorBlack)
	view.drawText(0, view.height-2, bartext2, termbox.ColorWhite, termbox.ColorBlack)
}

// debug line
func (view *View) drawDebugBar(text string) {
	view.drawText(0, viewHeight-1, text, termbox.ColorYellow, termbox.ColorBlack)
}

func (view *View) drawText(x int, y int, text string, fg termbox.Attribute, bg termbox.Attribute) {
	index := 0
	for _, char := range text {
		termbox.SetCell(x+index, y, rune(char), fg, bg)
		index++
	}
}
