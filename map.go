/*
a list of floors
*/
package main

func initMap() {
	gameMap = GameMap{
		floorMap:  make(FloorMap),
		floorList: FloorList{"0", "1", "2"}}

	for _, name := range gameMap.floorList {
		gameMap.floorMap[name] = generateFloor(name)
	}
}

func placePlayer(floorid string, tile rune) {
	floor := gameMap.floorMap[floorid]
	// positioning the player to stairs down
	for i, row := range floor {
		for j := range row {
			if floor[i][j].TileType == tile {
				player.PosX = i
				player.PosY = j
				return
			}
		}
	}
}

func switchFloor(operation string) {
	var toFloor string
	switch operation {
	case "+1":
		for index, name := range gameMap.floorList {
			if name == player.currentFloor {
				toFloor = gameMap.floorList[index+1]
				break
			}
		}
        player.currentFloor = toFloor
        placePlayer(toFloor, '<')
	case "-1":
		for index, name := range gameMap.floorList {
			if name == player.currentFloor {
				toFloor = gameMap.floorList[index-1]
				break
			}
		}
		player.currentFloor = toFloor
		placePlayer(toFloor, '>')
	}
}
