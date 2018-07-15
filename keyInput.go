/*
handling all player input
*/

package main

import (
	"github.com/nsf/termbox-go"
	"os"
)

func (keyInput *KeyInput) handleInput() {

	switch event := termbox.PollEvent(); event.Type {
	case termbox.EventKey:

		// player movement
		switch {
		case event.Key == termbox.KeyArrowUp || event.Ch == '8':
			player.move("N")
		case event.Ch == '9':
			player.move("NE")
		case event.Key == termbox.KeyArrowRight || event.Ch == '6':
			player.move("E")
		case event.Ch == '3':
			player.move("SE")
		case event.Key == termbox.KeyArrowDown || event.Ch == '2':
			player.move("S")
		case event.Ch == '1':
			player.move("SW")
		case event.Key == termbox.KeyArrowLeft || event.Ch == '4':
			player.move("W")
		case event.Ch == '7':
			player.move("NW")
		}

		// exit keys
		if event.Key == termbox.KeyEsc || event.Ch == 'q' {
			view.stop()
			os.Exit(0)
		}
	}
}
