/*
helper functions
*/

package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
	"os"
	"time"
)

// generating a random int [0,n)
func randomInt(upper int) int {
	randseed := rand.NewSource(time.Now().UnixNano())
	randsrc := rand.New(randseed)

	return randsrc.Intn(upper)
}

//directions on the map
func bydir(xpos int, ypos int, dir string, by int) (int, int) {
	switch dir {
	case "N":
		return xpos, ypos - by
	case "NE":
		return xpos + by, ypos - by
	case "E":
		return xpos + by, ypos
	case "SE":
		return xpos + by, ypos + by
	case "S":
		return xpos, ypos + by
	case "SW":
		return xpos - by, ypos + by
	case "W":
		return xpos - by, ypos
	case "NW":
		return xpos - by, ypos - by
	default:
		error("helperfuncs bydir(): unknown direction!")
	}
	return 0, 0
}

// crash with an error
func error(text string) {
	termbox.Close()
	fmt.Printf(text + "\n")
	os.Exit(-1)
}
