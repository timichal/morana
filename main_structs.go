/*
main structures and instantiations used all over
*/

package main

type (
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
		Name    string
		Level   int
		HP      int
		Attack  int
		Defense int
		Floor   string
		xpos    int
		ypos    int
	}
)

var (
	moves int
	view     View
	keyInput KeyInput
	player   = Player{Name: "Bob", Level: 1, HP: 100, Floor: "0"}
	tileset  = make(Tileset)
	floormap = make(FloorMap)
)
