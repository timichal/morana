/*
Main file of the Morana game
*/
package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	for i := 0; i<= 100; i++ {
		termbox.SetCell(i, 0, 'x', termbox.ColorGreen, termbox.ColorBlack)
		termbox.SetCell(i, 30, 'x', termbox.ColorGreen, termbox.ColorBlack)
		for j := 1; j < 30; j++ {
			termbox.SetCell(0, j, 'x', termbox.ColorGreen, termbox.ColorBlack)
			termbox.SetCell(100, j, 'x', termbox.ColorGreen, termbox.ColorBlack)
		}
	}
	termbox.Flush()
	MainLoop()
}

func MainLoop() {
	for {
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
