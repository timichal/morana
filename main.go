/*
Main file of the Morana game
*/
package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

var tiletypes = initTileTypes()

func main() {
	/*
		err := termbox.Init()
		if err != nil {
			panic(err)
		}

		defer termbox.Close()

		IntroScreen()
		MainLoop()
	*/

	fmt.Println(tiletypes)

	DrawMap()
}

func IntroScreen() {
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		termWidth, termHeight := termbox.Size()
		for i := 0; i < termWidth; i++ {
			termbox.SetCell(i, 0, 'x', termbox.ColorGreen, termbox.ColorBlack)
			termbox.SetCell(i, termHeight-1, 'x', termbox.ColorGreen, termbox.ColorBlack)
			for j := 1; j < termHeight-1; j++ {
				termbox.SetCell(0, j, 'x', termbox.ColorGreen, termbox.ColorBlack)
				termbox.SetCell(termWidth-1, j, 'x', termbox.ColorGreen, termbox.ColorBlack)
			}
		}
		termbox.Flush()

		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
			if event.Key == termbox.KeyEsc {
				return
			}
			if event.Ch == 'q' {
				return
			}
		}
	}
}

func MainLoop() {
}

func DrawMap() {

	//generating the map
	floormap := mapgen()
	player := placePlayer(floormap)

	for x, row := range floormap {
		for y, tile := range row {
			if x == player.xpos && y == player.ypos {
				fmt.Printf("@")
			} else {
				fmt.Printf(string(tile.TileType))
			}
		}
		fmt.Printf("\n")
	}

}
