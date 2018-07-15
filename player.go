/*
the player struct and methods
*/
package main

// positioning the player to stairs down
func (player *Player) init(floor Floor) {
	player.xpos = 10
	player.ypos = 10

	for i, row := range floor {
		for j := range row {
			if floor[i][j].TileType == '<' {
				player.xpos = i
				player.ypos = j
				return
			}
		}
	}
}

func (player *Player) move(dir string) {
	NewPosX, NewPosY := bydir(player.xpos, player.ypos, dir, 1)
	if tileset[floormap[player.Floor][NewPosX][NewPosY].TileType].Passable {
		player.xpos, player.ypos = NewPosX, NewPosY
	}
	moves++
	floorevents()
}

func (player *Player) changeHP(chval int) {
	player.HP = player.HP + chval
}
