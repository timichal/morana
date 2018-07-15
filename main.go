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

	gameInit()
	initView()
	// main loop
	for {
		view.refresh()
	}
}

func gameInit() {
	initEngine()
	initTileset()
	initMap()
	initPlayer()
	player.position(floormap["0"])
}
