/*
stuff that happens to the player on the floor
*/

package main

import (
//"os"
)

func floorevents() {
	switch gameMap.floorMap[player.currentFloor][player.PosX][player.PosY].TileType {
	case 'V':
		engine.State = "Victory"
	case '>':
		switchFloor("+1")
	case '<':
		switchFloor("-1")
	}
}
