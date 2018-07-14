/*
"objects" - types and structs
*/
package main

type (
	Player struct {
		Name    string
		Level   int
		Attack  int
		Defense int
		HP      int
	}

	Monster struct {
		Name    string
		Attack  int
		Defense int
		HP      int
	}
)

func (player *Player) changePlayerHP(chval int) {
	player.HP = player.HP + chval
}

func (player Player) changeDefPlayerHP(chval int) {
	player.HP = player.HP + chval
}
