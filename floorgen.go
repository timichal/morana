/*
floor generator
floor is 50x20 with . as empty spaces, # as wall
*/
package main

func floorgen() Floor {
	// initializing an empty 2d field
	var floor Floor

	// filling the filed with empty tiles
	for i, row := range floor {
		for j := range row {
			floor[i][j] = floor.genTile('.')
		}
	}

	// walls on the edges
	for pos := range floor[0] {
		floor[0][pos] = floor.genTile('#')
		floor[len(floor)-1][pos] = floor.genTile('#')
	}

	for pos := range floor {
		floor[pos][0] = floor.genTile('#')
		floor[pos][len(floor[0])-1] = floor.genTile('#')
	}

	// stairs down (for now)
	for {
		stairsx := randomInt(50)
		stairsy := randomInt(20)

		if floor[stairsx][stairsy].TileType == '.' {
			floor[stairsx][stairsy] = floor.genTile('<')
			break
		}
	}

	// victory flag for the lulz
	for {
		flagx := randomInt(50)
		flagy := randomInt(20)

		if floor[flagx][flagy].TileType == '.' {
			floor[flagx][flagy] = floor.genTile('V')
			break
		}
	}

	return floor
}

func (floor *Floor) genTile(tiletype rune) Tile {
	return Tile{TileType: tiletype}
}
