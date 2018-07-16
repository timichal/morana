/*
handling all player input
*/

package main

import (
	"github.com/nsf/termbox-go"
	"os"
)

func NewKeyInput() KeyInput {
	return KeyInput{
		chanKeyInput: make(chan *termbox.Event, 8),
	}
}

func (keyInput *KeyInput) Run() {
	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey && len(keyInput.chanKeyInput) < 8 {
			keyInput.chanKeyInput <- &event
		}
	}
}

func (keyInput *KeyInput) ProcessEvent(event *termbox.Event) {
	switch {
	case engine.State == "GameOn":
		// player movement
		switch {
		case event.Key == termbox.KeyArrowUp || event.Ch == '8':
			player.move("N")
		case event.Ch == '9':
			player.move("NE")
		case event.Key == termbox.KeyArrowRight || event.Ch == '6':
			player.move("E")
		case event.Ch == '3':
			player.move("SE")
		case event.Key == termbox.KeyArrowDown || event.Ch == '2':
			player.move("S")
		case event.Ch == '1':
			player.move("SW")
		case event.Key == termbox.KeyArrowLeft || event.Ch == '4':
			player.move("W")
		case event.Ch == '7':
			player.move("NW")
		}
	case engine.State == "Intro":
		engine.State = "GameOn"
	}

	// exit keys
	if event.Key == termbox.KeyEsc || event.Ch == 'q' {
		view.stop()
		os.Exit(0)
	}

	if event.Ch == 'r' {
		gameInit()
	}
}
