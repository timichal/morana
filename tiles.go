/*
a list of tile types
*/
package main

type (
	TileType struct {
		Name     string
		Passable bool
	}
)

func initTileTypes() map[rune]TileType {
	tiletypes := make(map[rune]TileType)

	//basic empty floor
	tiletypes['.'] = TileType{
		Name:     "Floor",
		Passable: true}

	//wall
	tiletypes['#'] = TileType{
		Name:     "Wall",
		Passable: false}

	return tiletypes
}
