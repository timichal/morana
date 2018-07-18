/*
Main file of the Morana game
*/
package main

import (
//"fmt"
//"github.com/nsf/termbox-go"
)

// main() tailored to termbox
func main() {
	initEngine()
	initTileset()
	engine.State = "Intro"
	initMap()
	initPlayer()
	player.position(floormap["0"])
	initView()
	view.run()
	// main loop
	engine.run()
	view.stop()
}

func restart() {
	engine.State = "Intro"
	initMap()
	initPlayer()
	player.position(floormap["0"])
	view.refresh()
}
