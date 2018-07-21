/*
a list of tile types
*/
package main

func initTileset() {
	//basic empty floor
	tileset['.'] = TileDef{
		Name:     "Floor",
		Passable: true}

	//wall
	tileset['#'] = TileDef{
		Name:     "Wall",
		Passable: false}

	//tower entrance/exit
	tileset['E'] = TileDef{
		Name:     "Entrance",
		Passable: true}

	//stairs down
	tileset['<'] = TileDef{
		Name:     "Stairs down",
		Passable: true}

	//stairs up
	tileset['>'] = TileDef{
		Name:     "Stairs up",
		Passable: true}

	//victory flag
	tileset['V'] = TileDef{
		Name:     "Victory",
		Passable: true}
}
