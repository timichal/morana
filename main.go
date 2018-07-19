/*
Main file of the Morana game
*/
package main

import (
	"log"
	"os"
	//"fmt"
	//"github.com/nsf/termbox-go"
)

// main() tailored to termbox
func main() {
	initLogger()
	initEngine()
	initTileset()
	initView()
	initGame()
	// main loop
	engine.run()
	view.stop()
}

func initGame() {
	engine.State = "Intro"
	initMap()
	initPlayer()
	placePlayer(player.currentFloor, 'E')
	view.run()
}

// logging
func initLogger() {
	logger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.LUTC|log.Llongfile)
	logFile, err := os.OpenFile("./morana.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("error opening logFile:", err)
	}
	defer logFile.Close()
	logger.SetOutput(logFile)
}
