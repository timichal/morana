/*
a list of floors
*/
package main

func initMap() {
	// level list
	gameMap.progression = []string{
		"0",
		"1",
		"2"}
	gameMap.floorIsGen = make([]bool, len(gameMap.progression))

	for i, _ := range gameMap.progression {
		gameMap.floorIsGen[i] = false
	}
}

func (gameMap *GameMap) getFloorIndex(floorName string) (floorIndex int) {
	for i, floor := range gameMap.progression {
		if floor == floorName {
			return i
		}
	}
	return -1
}
