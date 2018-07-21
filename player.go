/*
the player struct and methods
*/
package main

func initPlayer() {
	player = Player{
		Name:         "Bob",
		Level:        1,
		HP:           100}

	player.setFloor("init")
}

func (player *Player) move(dir string) {
	NewPosX, NewPosY := bydir(player.coord.X, player.coord.Y, dir, 1)
	if (NewPosX >= 0) && (NewPosX < floorWidth) && (0 <= NewPosY) && (NewPosY < floorHeight) {
		if tileset[gameMap.floorSet[player.currentFloor][NewPosX][NewPosY].TileType].Passable {
			player.coord.X, player.coord.Y = NewPosX, NewPosY
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

	// on-the-fly floor generation
	if !gameMap.floorIsGen[toFloor] {
		gameMap.floorSet[toFloor] = gameMap.generateFloor(toFloor)
		gameMap.floorIsGen[toFloor] = true
	}

    player.currentFloor = toFloor
    placePlayer(toFloor, placeMarker)
}
	