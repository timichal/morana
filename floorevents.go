/*
stuff that happens to the player on the floor
*/

package main

import (
//"os"
)

func floorevents() {
	switch gameMap.floorSet[player.currentFloor][player.coord.X][player.coord.Y].TileType {
	case 'V':
		engine.State = "Victory"
	case '>':
		player.setFloor("+1")
	case '<':
		player.setFloor("-1")
	}
}
