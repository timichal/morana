/*
initial attempt at a static single-level map
map is 20x50 with . as empty spaces, # as wall
*/
package main

import (
//"fmt"
)

type (
	Tile struct {
		//empty / wall
		TileType rune
		Explored bool
		Content  struct {
			Player bool
		}
	}
)

func GenEmptyTile() Tile {
	return Tile{TileType: '.'}
}

func GenWallTile() Tile {
	return Tile{TileType: '#'}
}

// generating an empty map
func mapgen() [20][50]Tile {
	mymap := [20][50]Tile{}

	//empty tiles first
	for i, row := range mymap {
		for j, _ := range row {
			mymap[i][j] = GenEmptyTile()
		}
	}

	//walls on the edges
	for pos, _ := range mymap[0] {
		mymap[0][pos] = GenWallTile()
		mymap[len(mymap)-1][pos] = GenWallTile()
	}

	for pos, _ := range mymap {
		mymap[pos][0] = GenWallTile()
		mymap[pos][len(mymap[0])-1] = GenWallTile()
	}

	return mymap
}
