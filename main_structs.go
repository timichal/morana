/*
main structures and instantiations used all over
*/

package main

import "github.com/nsf/termbox-go"

const (
	viewWidth   = 80
	viewHeight  = 24
	floorWidth  = 80
	floorHeight = 20
)

type (
	Engine struct {
		chanStop chan struct{}
		// engine states: Intro, GameOn, Victory
		State    string
		KeyInput KeyInput
	}

	View struct {
		width  int
		height int
	}

	KeyInput struct {
		chanKeyInput chan *termbox.Event
	}

	Tileset map[rune]TileType

	TileType struct {
		Name     string
		Passable bool
	}

	FloorMap map[string]Floor

	Floor [floorWidth][floorHeight]Tile

	Tile struct {
		//empty: . / wall: #
		TileType rune
		Explored bool
		Content  struct {
			Player bool
		}
	}

	Player struct {
		Name         string
		Level        int
		HP           int
		Attack       int
		PosX         int
		PosY         int
		Moves        int
		CurrentFloor string
	}
)

var (
	view      View
	keyInput  KeyInput
	engine    Engine
	player    Player
	debugText string
	tileset   = make(Tileset)
	floormap  = make(FloorMap)
)
