/*
Main file of the Morana game
*/
package main

import (
	//"fmt"
	"github.com/nsf/termbox-go"
)

var tiletypes = initTileTypes()
var floormap = mapgen()
var player = placePlayer(floormap)

func main() {

	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	//fmt.Println(tiletypes)

	drawMap()
}

func MainLoop() {
}

/*
func DrawMap() {
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
*/
