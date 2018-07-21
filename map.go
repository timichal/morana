/*
a list of floors
*/
package main

func initMap() {
	gameMap = GameMap{
		floorSet:    make(map[string]Floor),
        floorIsGen:  make(map[string]bool),
		progression: []string{"0", "1", "2"}}
}
