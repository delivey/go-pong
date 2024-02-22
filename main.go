package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var SCREEN_WIDTH = 320
var SCREEN_HEIGHT = 240

type Position struct {
	X, Y float32
}

type Ball struct {
	Position
	Radius float32
}

type Game struct {
	Ball        Ball
	CurrentTurn int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) HandleBall(screen *ebiten.Image) {
	g.Ball.X += 2.5
	if (g.Ball.X) > float32(SCREEN_WIDTH) {
		g.Ball.X = float32(SCREEN_WIDTH) / 2
	}
	vector.DrawFilledCircle(screen, g.Ball.X, g.Ball.Y, g.Ball.Radius, color.RGBA{255, 255, 255, 255}, false)
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Left player
	vector.DrawFilledRect(screen, 20, 50, 6, 30, color.RGBA{255, 255, 255, 255}, true)

	g.HandleBall(screen)

	// Right player
	vector.DrawFilledRect(screen, 294, 50, 6, 30, color.RGBA{255, 255, 255, 255}, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func (g *Game) Init() {
	g.Ball.X = float32(SCREEN_WIDTH) / 2
	g.Ball.Y = float32(SCREEN_HEIGHT) / 2
	g.Ball.Radius = 3
}

func main() {
	ebiten.SetWindowSize(960, 720)
	ebiten.SetWindowTitle("Hello, World!")
	game := &Game{}
	game.Init()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
