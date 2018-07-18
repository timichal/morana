/*
floor generator
floor is 80x20 with . as empty spaces, # as wall
*/
package main

import (
	"fmt"
	//"strconv"
)

//initial work - generate some rooms & use a pathfinding algorithm to connect them
func floorgenNew() Floor {
	var floor Floor
	var floorgen FloorGen
	var rooms Rooms

	// filling the floor with walls
	floor = floorgen.fillFloor(floor)

	// generating rooms for n attempts
	roomAttempts := 0
	for roomAttempts < maxRoomAttempts {
		var room Room
		room = floorgen.genRoom()
		if (room.topLeftCoord != Coord{-1, -1}) {
			rooms = append(rooms, room)
		}
		roomAttempts++
	}

	// check for overlap/closeness
	for i, room := range rooms {
		// could be improved: check just the unchecked part of the range
		for _, cmpRoom := range rooms {
			if room != cmpRoom {
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
		roomPathCoordX := room.topLeftCoord.X + randomInt(room.width)
		roomPathCoordY := room.topLeftCoord.Y + randomInt(room.height)
		rooms[ind].pathCoord = Coord{roomPathCoordX, roomPathCoordY}
	}

	debugText = fmt.Sprintln(rooms)
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
	roomTopLeftX, roomTopLeftY := randomInt(floorWidth-minRoomWidth), randomInt(floorHeight-minRoomHeight)

	roomWidth := randomInt(maxRoomWidth-minRoomWidth+1) + minRoomWidth
	roomHeight := randomInt(maxRoomHeight-minRoomHeight+1) + minRoomHeight

	if (roomTopLeftX+roomWidth <= floorWidth) && (roomTopLeftY+roomHeight <= floorHeight) {
		return Room{Coord{roomTopLeftX, roomTopLeftY}, Coord{-1, -1}, roomWidth, roomHeight, true}
	} else {
		return Room{topLeftCoord: Coord{-1, -1}}
	}
}

func (floorgen *FloorGen) genTile(tiletype rune) Tile {
	return Tile{TileType: tiletype}
}
