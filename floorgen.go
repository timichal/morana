/*
floor generator
floor is 80x20 with . as empty spaces, # as wall
*/
package main

import (
//"fmt"
//"strconv"
)

//initial work - generate some rooms & use a pathfinding algorithm to connect them
func floorgenNew() Floor {
	var floor Floor
	var floorgen FloorGen

	// filling the map with walls
	floor = floorgen.fillFloor(floor)

	// carving the rooms
	roomAttempts := 0
	for roomAttempts < maxRoomAttempts {
		floor = floorgen.makeRoom(floor)
		roomAttempts++
	}

	// stairs down (for now)
	for {
		stairsx := randomInt(floorWidth)
		stairsy := randomInt(floorHeight)

		if floor[stairsx][stairsy].TileType == '.' {
			floor[stairsx][stairsy] = floorgen.genTile('<')
			break
		}
	}

	return floor
}

func (floorgen *FloorGen) fillFloor(floor Floor) Floor {
	for i, row := range floor {
		for j := range row {
			floor[i][j] = floorgen.genTile('#')
		}
	}

	return floor
}

func (floorgen *FloorGen) makeRoom(floor Floor) Floor {
	// carving a room
	roomTopLeftX, roomTopLeftY := randomInt(floorWidth-minRoomWidth), randomInt(floorHeight-minRoomHeight)

	roomWidth := randomInt(maxRoomWidth-minRoomWidth+1) + minRoomWidth
	roomHeight := randomInt(maxRoomHeight-minRoomHeight+1) + minRoomHeight
	//debugText = fmt.Sprintln("mapx: " + strconv.Itoa(roomTopLeftX) + " mapy: " + strconv.Itoa(roomTopLeftY) + " width: " + strconv.Itoa(roomWidth) + " height: " + strconv.Itoa(roomHeight))

	// boundary check
	if (roomTopLeftX+roomWidth <= floorWidth) && (roomTopLeftY+roomHeight <= floorHeight) {

		// overlap/spacing check
		overlap := false

		// careful with map boundaries
		roomLowerBoundX := max(0, roomTopLeftX-minRoomSpacing)
		roomLowerBoundY := max(0, roomTopLeftY-minRoomSpacing)

		// *2 to compensate for lower bound
		roomUpperBoundX := roomWidth + minRoomSpacing*2
		roomUpperBoundY := roomHeight + minRoomSpacing*2
		//the check itself
	checkLoop:
		for i := 0; i < roomUpperBoundX; i++ {
			for j := 0; j < roomUpperBoundY; j++ {
				
				roomCheckX := min(roomLowerBoundX+i, floorWidth-1)
				roomCheckY := min(roomLowerBoundY+j, floorHeight-1)
				if floor[roomCheckX][roomCheckY].TileType == '.' {
					overlap = true
					break checkLoop
				}
			}
		}

		// placing the room
		if !overlap {
			for i := 0; i < roomWidth; i++ {
				for j := 0; j < roomHeight; j++ {
					floor[roomTopLeftX+i][roomTopLeftY+j] = floorgen.genTile('.')
				}
			}
		}
	}
	return floor
}

func (floorgen *FloorGen) genTile(tiletype rune) Tile {
	return Tile{TileType: tiletype}
}