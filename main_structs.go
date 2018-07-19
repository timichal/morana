/*
main structures and instantiations used all over
*/

package main

import (
	"github.com/nsf/termbox-go"
	"log"
	"math/rand"
	"time"
)

const (
	viewWidth   = 80
	viewHeight  = 24
	floorWidth  = 80
	floorHeight = 20

	minRoomWidth = 4
	maxRoomWidth = 15

	minRoomHeight = 3
	maxRoomHeight = 6

	maxRoomAttempts = 30
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

	PathCoords []Coord

	Coord struct {
		X int
		Y int
	}

	KeyInput struct {
		chanKeyInput chan *termbox.Event
	}

	Tileset map[rune]TileType

	TileType struct {
		Name     string
		Passable bool
	}

	GameMap struct {
		floorList FloorList
		floorMap  FloorMap
	}

	FloorList []string

	FloorMap map[string]Floor

	FloorGen struct {
		floorName string
		floor     Floor
		rooms     Rooms
	}

	Floor [floorWidth][floorHeight]Tile

	Rooms []Room

	Room struct {
		topLeftCoord Coord
		pathCoord    Coord
		width        int
		height       int
		isValid      bool
	}

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
		currentFloor string
	}
)

var (
	randseed  = rand.NewSource(time.Now().UnixNano())
	view      View
	keyInput  KeyInput
	engine    Engine
	player    Player
	gameMap   GameMap
	floorgen  FloorGen
	debugText string
	logger    *log.Logger
	tileset   = make(Tileset)
)
