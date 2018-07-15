/*
the game engine
*/

package main

import "github.com/nsf/termbox-go"

func initEngine() {
	engine = Engine{
		chanStop: make(chan struct{}, 1),
		// engine states: Intro, GameOn, Victory
		State:        "Intro",
		CurrentFloor: "0"}
}

func (engine *Engine) run() {
	engine.KeyInput = NewKeyInput()
	go engine.KeyInput.Run()

	var event *termbox.Event

	for {
		select {
		case <-engine.chanStop:
			break 
		case event = <-engine.KeyInput.chanKeyInput:
			engine.KeyInput.ProcessEvent(event)
			view.refresh()

		}
	}
}	