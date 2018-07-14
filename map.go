/*
initial attempt at a static single-level map
map is 50x20 with . as empty spaces, # as wall
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
func mapgen() [50][20]Tile {
	mymap := [50][20]Tile{}

	//empty tiles first
	for i, row := range mymap {
		for j := range row {
			mymap[i][j] = GenEmptyTile()
		}
	}

	//walls on the edges
	for pos := range mymap[0] {
		mymap[0][pos] = GenWallTile()
		mymap[len(mymap)-1][pos] = GenWallTile()
	}

	for pos := range mymap {
		mymap[pos][0] = GenWallTile()
		mymap[pos][len(mymap[0])-1] = GenWallTile()
	}

	return mymap
}
