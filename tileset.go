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

	//tower entrance/exit
	tileset['E'] = TileType{
		Name:     "Entrance",
		Passable: true}

	//stairs down
	tileset['<'] = TileType{
		Name:     "Stairs down",
		Passable: true}

	//stairs up
	tileset['>'] = TileType{
		Name:     "Stairs up",
		Passable: true}

	//victory flag
	tileset['V'] = TileType{
		Name:     "Victory",
		Passable: true}
}
