/*
initial attempt at a static single-level map
map is 20x20 with . as empty spaces, # as wall
*/
package main

import "fmt"

type (
	Tile struct {
			//empty / wall
			TileType rune
			Explored bool
			Content struct {

			}
		}
)

func GenEmptyTile() Tile {
	return Tile{TileType:'.', Explored:false}
}

func mapgen() {
	tile := GenEmptyTile()
	fmt.Println(tile)
	/*
	//emptytile := Tile{TileType:'.', Explored:false}
	var row [20]Tile
	for i := 0; i < len(row); i++ {
		row[i].TileType = '.'
	}
	
	*/
}
