/*
the player struct and methods
*/
package main

func initPlayer() {
	player = Player{
		Name:         "Bob",
		currentFloor: "0",
		Level:        1,
		HP:           100}
}

func (player *Player) move(dir string) {
	NewPosX, NewPosY := bydir(player.PosX, player.PosY, dir, 1)
	if (NewPosX >= 0) && (NewPosX < floorWidth) && (0 <= NewPosY) && (NewPosY < floorHeight) {
		if tileset[gameMap.floorMap[player.currentFloor][NewPosX][NewPosY].TileType].Passable {
			player.PosX, player.PosY = NewPosX, NewPosY
			player.Moves++
			floorevents()
		}
	}

}

func (player *Player) changeHP(chval int) {
	player.HP = player.HP + chval
}
