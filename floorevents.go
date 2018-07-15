/*
stuff that happens to the player on the floor
*/

package main

import (
//"fmt"
//"os"
)

func floorevents() {
	//debugtext = fmt.Sprintln(player.PosX, player.PosY, floormap[player.Floor][player.PosX][player.PosY])
	// victory
	if floormap[player.Floor][player.PosX][player.PosY].TileType == 'V' {
		engine.State = "Victory"
	}
}
