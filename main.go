/*
Main file of the Morana game
*/
package main

import (
//"fmt"
//"github.com/nsf/termbox-go"
)
var debugtext string
// main() tailored to termbox
func main() {

	initTileset()
	initMap()

	player.init(floormap["0"])

	initView()

	// main loop
	for {
		view.refresh()
	}
}
