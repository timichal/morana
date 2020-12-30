package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

// Our game constants
const (
	screenWidth  = 240
	screenHeight = 240
)

const (
	tileSize = 16
	tileXNum = 25
)

// Create our empty vars
var (
	err        error
	ship       *ebiten.Image
	enemyImage *ebiten.Image
	tilesImage *ebiten.Image
	playerOne  player
	ticks      int
	canMove    bool
	dir        string
)

type Game struct {
	layers [][]int
}

// Create the player class
type player struct {
	image      *ebiten.Image
	xPos, yPos float64
	hp         float64
	attack     float64
}

type enemy struct {
	image      *ebiten.Image
	xPos, yPos float64
	hp         float64
	attack     float64
}

// Run this code once at startup
func init() {
	ship, _, err = ebitenutil.NewImageFromFile("assets/ship.png")
	if err != nil {
		log.Fatal(err)
	}

	enemyImage, _, err = ebitenutil.NewImageFromFile("assets/enemy.png")
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(bytes.NewReader(images.Tiles_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)

	playerOne = player{ship, 0, 0, 20, 1}

	dir = "right"
}

func generateEnemies() []enemy {
	enemies := make([]enemy, 1)
	enemies[0] = enemy{enemyImage, 2, 2, 5, 1}
	return enemies
}

func (g *Game) Update() error {
	ticks = (ticks + 1) % 60
	if ticks%20 == 0 {
		canMove = true
	}

	// move every second
	if ticks%20 == 0 {
		if dir == "right" {
			if playerOne.xPos < tileSize*14 {
				playerOne.xPos += tileSize
			} else {
				playerOne.yPos += tileSize
				dir = "left"
			}
		} else if dir == "left" {
			if playerOne.xPos > 0 {
				playerOne.xPos -= tileSize
			} else {
				playerOne.yPos += tileSize
				dir = "right"
			}
		}
	}
	/*
		if canMove == true {
			if ebiten.IsKeyPressed(ebiten.KeyUp) && playerOne.yPos > 0 {
				playerOne.yPos -= playerOne.speed
				canMove = false
			}
			if ebiten.IsKeyPressed(ebiten.KeyDown) && playerOne.yPos < tileSize*14 {
				playerOne.yPos += playerOne.speed
				canMove = false
			}
			if ebiten.IsKeyPressed(ebiten.KeyLeft) && playerOne.xPos > 0 {
				playerOne.xPos -= playerOne.speed
				canMove = false
			}
			if ebiten.IsKeyPressed(ebiten.KeyRight) && playerOne.xPos < tileSize*14 {
				playerOne.xPos += playerOne.speed
				canMove = false
			}
		}
	*/
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw each tile with each DrawImage call.
	// As the source images of all DrawImage calls are always same,
	// this rendering is done very effectively.
	// For more detail, see https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#Image.DrawImage
	const xNum = screenWidth / tileSize
	for _, l := range g.layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xNum)*tileSize), float64((i/xNum)*tileSize))

			sx := (t % tileXNum) * tileSize
			sy := (t / tileXNum) * tileSize
			screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f Tick: %d", ebiten.CurrentTPS(), ticks))

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(playerOne.xPos, playerOne.yPos)
	screen.DrawImage(playerOne.image, playerOp)

	enemies := generateEnemies()
	for _, enemy := range enemies {
		enemyOp := &ebiten.DrawImageOptions{}
		enemyOp.GeoM.Translate(enemy.xPos*tileSize, enemy.yPos*tileSize)
		screen.DrawImage(enemy.image, enemyOp)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (width, height int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{
		layers: [][]int{
			{
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 244, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,

				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 244, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 219, 243, 243, 243, 219, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,

				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 244, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			},
			/*
				{
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 26, 27, 28, 29, 30, 31, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 51, 52, 53, 54, 55, 56, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 76, 77, 78, 79, 80, 81, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 101, 102, 103, 104, 105, 106, 0, 0, 0, 0,

					0, 0, 0, 0, 0, 126, 127, 128, 129, 130, 131, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 303, 303, 245, 242, 303, 303, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,

					0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
				},
			*/
		},
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Tiles (Ebiten Demo)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
