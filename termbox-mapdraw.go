/*
drawing the map with the termbox library
*/

package main

import (
	"github.com/nsf/termbox-go"
)

func drawMap() {
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		
		//termWidth, termHeight := 50, 20

		// draw map by tile type
		for x, row := range floormap {
			for y, tile := range row {
				termbox.SetCell(x, y, tile.TileType, termbox.ColorGreen, termbox.ColorBlack)
			}
		}

		// draw the player
		termbox.SetCell(player.xpos, player.ypos, '@', termbox.ColorRed, termbox.ColorBlack)

		termbox.Flush()

		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
			//movement keys
			if event.Key == termbox.KeyArrowLeft {
				player.xpos = player.xpos - 1
			} else if event.Key == termbox.KeyArrowRight {
				player.xpos = player.xpos + 1
			} else if event.Key == termbox.KeyArrowUp {
				player.ypos = player.ypos - 1
			} else if event.Key == termbox.KeyArrowDown {
				player.ypos = player.ypos + 1
			}

			// exit keys
			if event.Key == termbox.KeyEsc {
				return
			}
			if event.Ch == 'q' {
				return
			}
		}
	}
}
