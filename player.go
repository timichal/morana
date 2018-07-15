/*
the player struct and methods
*/
package main

func initPlayer() {
	player = Player{
		Name:  "Bob",
		Level: 1,
		HP:    100,
		Floor: "0"}
}

// positioning the player to stairs down
func (player *Player) position(floor Floor) {
	for i, row := range floor {
		for j := range row {
			if floor[i][j].TileType == '<' {
				player.PosX = i
				player.PosY = j
				return
			}
		}
	}
}

func (player *Player) move(dir string) {
	NewPosX, NewPosY := bydir(player.PosX, player.PosY, dir, 1)
	if tileset[floormap[player.Floor][NewPosX][NewPosY].TileType].Passable {
		player.PosX, player.PosY = NewPosX, NewPosY
		player.Moves++
		floorevents()
	}

}

func (player *Player) changeHP(chval int) {
	player.HP = player.HP + chval
}
