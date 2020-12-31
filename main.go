package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	player     Player
	ticks      int
	canMove    bool
	dir        bool
	enemies    []Enemy
)

type Game struct {
	layers [][]int
}

// Create the player class
type Player struct {
	image      *ebiten.Image
	xPos, yPos int
	hp         float64
	attack     float64
	exp        float64
}

type Enemy struct {
	image      *ebiten.Image
	xPos, yPos int
	hp         float64
	attack     float64
}

func generateEnemies(numberOfEnemies int) []Enemy {
	// generate coords
	coords := make([][2]int, numberOfEnemies)
	for i := range coords {
		coords[i] = [2]int{-1, -1}
	}

	for i := range coords {
		var x, y int
		for {
			x, y = rand.Intn(15), rand.Intn(15)

			duplicate := false
			for _, coord := range coords {
				if coord[0] == x && coord[1] == y {
					duplicate = true
				}
			}

			if duplicate == false {
				break
			}
		}

		coords[i] = [2]int{x, y}
	}

	// generate enemies
	enemies := make([]Enemy, numberOfEnemies)

	for i := range enemies {
		enemies[i] = Enemy{enemyImage, coords[i][0], coords[i][1], 5, 1}
	}

	return enemies
}

func reset() {
	rand.Seed(time.Now().UnixNano())
	dir = true // right
	player.xPos, player.yPos = 0, 0
	player.hp = 20
	enemies = generateEnemies(30)
}

// Run this code once at startup
func init() {
	ship, _, err = ebitenutil.NewImageFromFile("assets/hero.png")
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
	player = Player{ship, 0, 0, 20, 1, 0}
	reset()
}

func findEnemy(x, y int) int {
	for i, enemy := range enemies {
		if enemy.xPos == x && enemy.yPos == y {
			return i
		}
	}

	return -1
}

func step(nextX int, nextY int, reverseDir bool) {
	enemy := findEnemy(nextX, nextY)
	if enemy < 0 || enemies[enemy].hp <= 0 {
		player.xPos, player.yPos = nextX, nextY
		if reverseDir {
			dir = !dir
		}
	} else {
		fightEnemy(enemy)
	}
}

func fightEnemy(enemyIndex int) {
	enemy := &enemies[enemyIndex]
	fmt.Printf("Enemy HP: %f | Player HP: %f\n", enemy.hp, player.hp)
	// first the player attacks, then the enemy
	// todo speed
	enemy.hp -= player.attack
	if enemy.hp <= 0 {
		player.exp += 1
		return
	}

	player.hp -= enemy.attack
	if player.hp <= 0 {
		reset()
	}
}

func (g *Game) Update() error {
	ticks = (ticks + 1) % 60
	if ticks%20 == 0 {
		canMove = true
	}

	// move every second
	if ticks%20 == 0 {
		if dir {
			if player.xPos < 14 {
				step(player.xPos+1, player.yPos, false)
			} else {
				step(player.xPos, player.yPos+1, true)
			}
		} else {
			if player.xPos > 0 {
				step(player.xPos-1, player.yPos, false)
			} else {
				step(player.xPos, player.yPos+1, true)
			}
		}
	}

	// upgrade button
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x > 0 && x < 100 && y > 230 {
			if player.exp >= 1 {
				player.exp -= 1
				player.attack += 1
			}
		}
	}

	/*
		if canMove == true {
			if ebiten.IsKeyPressed(ebiten.KeyUp) && player.yPos > 0 {
				player.yPos -= player.speed
				canMove = false
			}
			if ebiten.IsKeyPressed(ebiten.KeyDown) && player.yPos < tileSize*14 {
				player.yPos += player.speed
				canMove = false
			}
			if ebiten.IsKeyPressed(ebiten.KeyLeft) && player.xPos > 0 {
				player.xPos -= player.speed
				canMove = false
			}
			if ebiten.IsKeyPressed(ebiten.KeyRight) && player.xPos < tileSize*14 {
				player.xPos += player.speed
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

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Tick: %d HP: %.0f XP: %.0f Atk: %.0f", ticks, player.hp, player.exp, player.attack))

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(float64(player.xPos)*tileSize, float64(player.yPos)*tileSize)
	screen.DrawImage(player.image, playerOp)

	for _, enemy := range enemies {
		if enemy.hp > 0 {
			enemyOp := &ebiten.DrawImageOptions{}
			enemyOp.GeoM.Translate(float64(enemy.xPos)*tileSize, float64(enemy.yPos)*tileSize)
			screen.DrawImage(enemy.image, enemyOp)
		}
	}

	drawTextWithShadow(screen, "upgrade atk", 0, 230, 1, color.Black)
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
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 129,
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
