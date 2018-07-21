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
	viewWidth    = 80
	viewHeight   = 24
	topBarOffset = 1
	floorWidth   = 80
	floorHeight  = 20

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

	Player struct {
		Name         string
		Level        int
		HP           int
		Attack       int
		coord        Coord
		moves        int
		currentFloor string
	}

	// monsters definitions: monsters.go
	MonsterSet map[string]Monster

	Monster struct {
		name    string
		group   rune
		HP      int
		attack  int
		defense int
	}

	// tiles definitions: tileset.go
	Tileset map[rune]TileDef

	TileDef struct {
		Name     string
		Passable bool
	}

	// map hierarchy: map.go & floorgen.go
	GameMap struct {
		progression []string
		floorSet    map[string]Floor
		floorIsGen map[string]bool
	}

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
		TileType rune
		Explored bool
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
	monsters  = make(MonsterSet)
)
