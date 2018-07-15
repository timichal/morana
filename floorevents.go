/*
stuff that happens to the player on the floor
*/

package main

import "fmt"
import "os"
import "github.com/nsf/termbox-go"
func floorevents() {
	// victory
	if floormap[player.Floor][player.xpos][player.ypos].TileType == '🏁'{
		termbox.Close()
		fmt.Printf("You win! \n")
		os.Exit(0)
	}
}