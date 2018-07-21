/*
monster definitions
*/

package main

func initMonsters() {
	// basic robot
	monsters["robot"] = Monster{
		name:    "Robot",
		group:   'R',
		HP:      10,
		attack:  5,
		defense: 5}
}
