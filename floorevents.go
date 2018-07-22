/*
stuff that happens to the player on the floor

called from plater.move()
*/

package main

import (
//"os"
)

func floorevents() {
	switch gameMap.floorSet[gameMap.getFloorIndex(player.currentFloor)].plan[player.coord.X][player.coord.Y].tileType {
	case 'V':
		engine.State = "Victory"
	case '>':
		player.setFloor("+1")
	case '<':
		player.setFloor("-1")
	}
}
