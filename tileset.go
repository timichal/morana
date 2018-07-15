/*
a list of tile types
*/
package main

func initTileset() {
	//basic empty floor
	tileset['.'] = TileType{
		Name:     "Floor",
		Passable: true}

	//wall
	tileset['#'] = TileType{
		Name:     "Wall",
		Passable: false}

	//stairs down
	tileset['<'] = TileType{
		Name:     "Stairs down",
		Passable: true}

	//victory flag
	tileset['🏁'] = TileType{
		Name:     "Victory",
		Passable: true}
}
