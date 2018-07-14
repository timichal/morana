/*
Main file of the Morana game
*/
package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func main() {
	/*
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	IntroScreen()
	MainLoop()
	*/
	playertest := Player{Name: "Bob", Level: 1, Attack: 10, Defense: 10, HP: 100}
	fmt.Printf("%d\n", playertest.HP)
	playertest.changePlayerHP(10)
	fmt.Printf("%d\n", playertest.HP)

	mapgen()
}

func IntroScreen() {
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		termWidth, termHeight := termbox.Size()
		for i := 0; i < termWidth; i++ {
			termbox.SetCell(i, 0, 'x', termbox.ColorGreen, termbox.ColorBlack)
			termbox.SetCell(i, termHeight-1, 'x', termbox.ColorGreen, termbox.ColorBlack)
			for j := 1; j < termHeight-1; j++ {
				termbox.SetCell(0, j, 'x', termbox.ColorGreen, termbox.ColorBlack)
				termbox.SetCell(termWidth-1, j, 'x', termbox.ColorGreen, termbox.ColorBlack)
			}
		}
		termbox.Flush()

		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
			if event.Key == termbox.KeyEsc {
				return
			}
			if event.Ch == 'q' {
				return
			}
		}
	}
}

func MainLoop() {
}
