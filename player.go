/*
the player struct and methods
*/
package main

func initPlayer() {
	player = Player{
		Name:  "Bob",
		Level: 1,
		HP:    100}

	player.setFloor("init")
}

func (player *Player) move(dir string) {
	NewPosX, NewPosY := bydir(player.coord.X, player.coord.Y, dir, 1)
	if (NewPosX >= 0) && (NewPosX < floorWidth) && (0 <= NewPosY) && (NewPosY < floorHeight) {
		if tileset[gameMap.floorSet[gameMap.getFloorIndex(player.currentFloor)].plan[NewPosX][NewPosY].tileType].Passable {
			gameMap.floorSet[gameMap.getFloorIndex(player.currentFloor)].plan[player.coord.X][player.coord.Y].content = []string{}
			player.coord.X, player.coord.Y = NewPosX, NewPosY
			gameMap.floorSet[gameMap.getFloorIndex(player.currentFloor)].plan[player.coord.X][player.coord.Y].content = append(gameMap.floorSet[gameMap.getFloorIndex(player.currentFloor)].plan[player.coord.X][player.coord.Y].content, "player")
			player.moves++
			floorevents()
		}
	}

}

func (player *Player) setFloor(operation string) {
	var toFloor string
	var placeMarker rune

	switch operation {
	case "init":
		toFloor = gameMap.progression[0]
		player.currentFloor = toFloor
		placeMarker = 'E'
	case "+1":
		for index, name := range gameMap.progression {
			if name == player.currentFloor {
				toFloor = gameMap.progression[index+1]
				break
			}
		}
		placeMarker = '<'
	case "-1":
		for index, name := range gameMap.progression {
			if name == player.currentFloor {
				toFloor = gameMap.progression[index-1]
				break
			}
		}
		placeMarker = '>'
	}
	toFloorIndex := gameMap.getFloorIndex(toFloor)
	// on-the-fly floor generation
	if !gameMap.floorIsGen[toFloorIndex] {
		gameMap.floorSet = append(gameMap.floorSet, FloorDesc{toFloor, gameMap.generateFloor(toFloor)})
		gameMap.floorIsGen[toFloorIndex] = true
	}

	player.currentFloor = toFloor
	player.placeToUniqTile(toFloor, placeMarker)
}

func (player *Player) placeToUniqTile(floorid string, tile rune) {
	floor := gameMap.floorSet[gameMap.getFloorIndex(floorid)].plan
	// positioning the player to stairs down
	for i, row := range floor {
		for j := range row {
			if floor[i][j].tileType == tile {
				gameMap.floorSet[gameMap.getFloorIndex(floorid)].plan[i][j].content = append(floor[i][j].content, "player")
				player.coord.X = i
				player.coord.Y = j
				return
			}
		}
	}
}
