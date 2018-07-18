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
	initView()
	initGame()
	// main loop
	engine.run()
	view.stop()
}

func initGame() {
	engine.State = "Intro"
	initMap()
	initPlayer()
	player.position(floormap["0"])
	view.run()
}
