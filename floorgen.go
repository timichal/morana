/*
floor generator
floor is 80x20 with . as empty spaces, # as wall
*/
package main

import (
	"fmt"
)

//initial work - generate some rooms & use a pathfinding algorithm to connect them
func (gameMap *GameMap) generateFloor(floorName string) Floor {
	var floorgen = FloorGen{
		floorName: floorName}

	// filling the floor with walls
	floorgen.fillFloor()

	// generating possible room attempts, capped at maxRoomAttempts)
	for roomAttempts := 0; roomAttempts < maxRoomAttempts; roomAttempts++ {
		floorgen.rooms = append(floorgen.rooms, floorgen.generateRoom())
	}

	// check for overlap
	for i, room := range floorgen.rooms {
		for j, cmpRoom := range floorgen.rooms {
			if j < i {
				if cmpRoom.isValid {
					if floorgen.isOverlapping(room, cmpRoom) {
						floorgen.rooms[i].isValid = false
						break
					}
				}
			}
		}
	}

	// filter out invalid rooms
	var validRooms Rooms
	for _, room := range floorgen.rooms {
		if room.isValid {
			validRooms = append(validRooms, room)
		}
	}
	floorgen.rooms = validRooms

	// carve the rooms into the map
	for ind, room := range floorgen.rooms {
		for i := 0; i < room.width; i++ {
			for j := 0; j < room.height; j++ {
				floorgen.floor[room.topLeftCoord.X+i][room.topLeftCoord.Y+j] = floorgen.genTile('.')
			}
		}

		// setting random point in room for pathfinding
		roomPathCoordX := room.topLeftCoord.X + randomInt(room.width)
		roomPathCoordY := room.topLeftCoord.Y + randomInt(room.height)
		floorgen.rooms[ind].pathCoord = Coord{roomPathCoordX, roomPathCoordY}
	}

	// rudimentary pathfinding
	floorgen.roomPathFind()

	// floor specific features
	floorgen.placeSpecifics()

	return floorgen.floor
}

// shamelessly stolen logic thingy
func (floorgen *FloorGen) isOverlapping(room, otherroom Room) bool {
	return (room.topLeftCoord.X <= otherroom.topLeftCoord.X+otherroom.width+1) &&
		(room.topLeftCoord.X+room.width >= otherroom.topLeftCoord.X-1) &&
		(room.topLeftCoord.Y <= otherroom.topLeftCoord.Y+otherroom.height+1) &&
		(room.topLeftCoord.Y+room.height >= otherroom.topLeftCoord.Y-1)
}

func (floorgen *FloorGen) fillFloor() {
	for i, row := range floorgen.floor {
		for j := range row {
			floorgen.floor[i][j] = floorgen.genTile('#')
		}
	}
}

func (floorgen *FloorGen) generateRoom() Room {
	for {
		roomTopLeftX, roomTopLeftY := randomInt(floorWidth-minRoomWidth), randomInt(floorHeight-minRoomHeight)

		roomWidth := randomInt(maxRoomWidth-minRoomWidth+1) + minRoomWidth
		roomHeight := randomInt(maxRoomHeight-minRoomHeight+1) + minRoomHeight

		if (roomTopLeftX+roomWidth <= floorWidth) && (roomTopLeftY+roomHeight <= floorHeight) {
			return Room{Coord{roomTopLeftX, roomTopLeftY}, Coord{-1, -1}, roomWidth, roomHeight, true}
		}
	}
}

func (floorgen *FloorGen) genTile(tiletype rune) Tile {
	return Tile{tileType: tiletype}
}

func (floorgen *FloorGen) roomPathFind() {
	for i, room := range floorgen.rooms {
		for j, cmpRoom := range floorgen.rooms {
			if j > i {
				debugText = fmt.Sprintln(room.pathCoord, cmpRoom.pathCoord)

				srcX := min(room.pathCoord.X, cmpRoom.pathCoord.X)
				var srcY, destX, destY int
				if srcX == room.pathCoord.X {
					destX = cmpRoom.pathCoord.X
					srcY = room.pathCoord.Y
					destY = cmpRoom.pathCoord.Y
				} else {
					destX = room.pathCoord.X
					srcY = cmpRoom.pathCoord.Y
					destY = room.pathCoord.Y
				}

				for srcX != destX {
					srcX++
					floorgen.floor[srcX][srcY] = floorgen.genTile('.')
				}

				for srcY != destY {
					if srcY < destY {
						srcY++
					} else {
						srcY--
					}
					floorgen.floor[srcX][srcY] = floorgen.genTile('.')
				}
				break
			}
		}
	}
}

func (floorgen *FloorGen) placeSpecifics() {
	floorname := floorgen.floorName
	//first floor
	if floorname == gameMap.progression[0] {
		// place entrance
		floorgen.placeRandom('E')
		//all except first floor
	} else {
		//place stairs down
		floorgen.placeRandom('<')
	}

	//last floor
	if floorname == gameMap.progression[len(gameMap.progression)-1] {
		// place victory flag
		floorgen.placeRandom('V')
		//all except last floor
	} else {
		// place stairs up
		floorgen.placeRandom('>')
	}
}

func (floorgen *FloorGen) placeRandom(tiletype rune) {
	for {
		x := randomInt(floorWidth)
		y := randomInt(floorHeight)

		if floorgen.floor[x][y].tileType == '.' {
			floorgen.floor[x][y] = floorgen.genTile(tiletype)
			break
		}
	}
}
