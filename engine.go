/*
the game engine
*/

package main

import "github.com/nsf/termbox-go"

func initEngine() {
	engine = Engine{
		chanStop: make(chan struct{}, 1)}
}

func (engine *Engine) run() {
	engine.KeyInput = NewKeyInput()
	go engine.KeyInput.Run()

	var event *termbox.Event

loop:
	for {
		select {
		case <-engine.chanStop:
			break loop
		default:
			select {
			case event = <-engine.KeyInput.chanKeyInput:
				engine.KeyInput.ProcessEvent(event)
				view.refresh()
			case <-engine.chanStop:
				break loop
			}
		}
	}
}

func (engine *Engine) Stop() {
	close(engine.chanStop)
}
