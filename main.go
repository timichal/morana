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
	initMap()
	initPlayer()
	player.position(floormap["0"])
	initView()
	// main loop
	engine.run()
	view.stop()
}

func gameInit() {

}
