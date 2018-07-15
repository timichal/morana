/*
main structures and instantiations used all over
*/

package main

type (
	Engine struct {
		Victory bool
	}

	View struct {
	}

	KeyInput struct {
	}

	Tileset map[rune]TileType

	TileType struct {
		Name     string
		Passable bool
	}

	FloorMap map[string]Floor

	Floor [50][20]Tile

	Tile struct {
		//empty: . / wall: #
		TileType rune
		Explored bool
		Content  struct {
			Player bool
		}
	}

	Player struct {
		Name   string
		Level  int
		HP     int
		Attack int
		Floor  string
		PosX   int
		PosY   int
		Moves  int
	}
)

var (
	view     View
	keyInput KeyInput
	engine   Engine
	player   Player
	tileset  = make(Tileset)
	floormap = make(FloorMap)
)
