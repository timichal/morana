/*
the player struct and methods
*/
package main

type (
	Player struct {
		Name    string
		Level   int
		Attack  int
		Defense int
		HP      int
		xpos    int
		ypos    int
	}
)

func (player *Player) changePlayerHP(chval int) {
	player.HP = player.HP + chval
}

func placePlayer(floormap [50][20]Tile) Player {
	for {
		xpos := randomInt(50)
		ypos := randomInt(20)

		tiletype := floormap[xpos][ypos].TileType

		if tiletypes[tiletype].Passable {
			player := Player{xpos: xpos, ypos: ypos}
			return player
		}
	}
}
