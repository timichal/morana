/*
floor generator
floor is 80x20 with . as empty spaces, # as wall
*/
package main

import (
	"fmt"
)

//initial work - generate some rooms & use a pathfinding algorithm to connect them
func floorgenNew() Floor {
	var floor Floor
	var floorgen FloorGen
	var rooms Rooms

	// filling the floor with walls
	floor = floorgen.fillFloor(floor)

	// generating rooms for n attempts
	for roomAttempts := 0; roomAttempts < maxRoomAttempts; roomAttempts++ {
		rooms = append(rooms, floorgen.genRoom())
	}

	// check for overlap/closeness
	for i, room := range rooms {
		for j, cmpRoom := range rooms {
			if j > i {
				if cmpRoom.isValid {
					if floorgen.isOverlapping(room, cmpRoom) {
						rooms[i].isValid = false
						break
					}
				}
			}
		}
	}

	// filter out invalid rooms
	var validRooms Rooms
	for _, room := range rooms {
		if room.isValid {
			validRooms = append(validRooms, room)
		}
	}
	rooms = validRooms

	// draw the rooms on the map
	for ind, room := range rooms {
		for i := 0; i < room.width; i++ {
			for j := 0; j < room.height; j++ {
				floor[room.topLeftCoord.X+i][room.topLeftCoord.Y+j] = floorgen.genTile('.')
			}
		}

		// has to be an edge coord!
		roomPathCoordX := room.topLeftCoord.X + randomInt(room.width)
		roomPathCoordY := room.topLeftCoord.Y + randomInt(room.height)
		rooms[ind].pathCoord = Coord{roomPathCoordX, roomPathCoordY}
	}

	debugText = fmt.Sprintln(rooms)

	// rudimentary pathfinding
	for i, room := range rooms {
		for j, cmpRoom := range rooms {
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
					floor[srcX][srcY] = floorgen.genTile('.')
				}
				
				for srcY != destY {
					if srcY < destY {
						srcY++
					} else {
						srcY--
					}
					floor[srcX][srcY] = floorgen.genTile('.')
				}				
				break
			}
		}
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

	// v-flag
	for {
		victx := randomInt(floorWidth)
		victy := randomInt(floorHeight)

		if floor[victx][victy].TileType == '.' {
			floor[victx][victy] = floorgen.genTile('V')
			break
		}
	}

	return floor
}

// shamelessly stolen logic thingy
func (floorgen *FloorGen) isOverlapping(room, otherroom Room) bool {
	return (room.topLeftCoord.X <= otherroom.topLeftCoord.X+otherroom.width+1) &&
		(room.topLeftCoord.X+room.width >= otherroom.topLeftCoord.X-1) &&
		(room.topLeftCoord.Y <= otherroom.topLeftCoord.Y+otherroom.height+1) &&
		(room.topLeftCoord.Y+room.height >= otherroom.topLeftCoord.Y-1)
}

func (floorgen *FloorGen) fillFloor(floor Floor) Floor {
	for i, row := range floor {
		for j := range row {
			floor[i][j] = floorgen.genTile('#')
		}
	}
	return floor
}

func (floorgen *FloorGen) genRoom() Room {
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
	return Tile{TileType: tiletype}
}
