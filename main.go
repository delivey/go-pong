package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

type Paddle struct {
	Position
	Height float32
	Width  float32
}

type Game struct {
	Ball         Ball
	IsPlayerTurn bool
	BallSpeed    float32
	Player       Paddle
	Bot          Paddle
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) HandleBall(screen *ebiten.Image) {
	switch g.IsPlayerTurn {
	case true:
		g.Ball.X -= g.BallSpeed
	case false:
		g.Ball.X += g.BallSpeed
	}
	if g.Ball.X > float32(SCREEN_WIDTH) || g.Ball.X < 0 {
		g.Ball.X = float32(SCREEN_WIDTH) / 2
		g.IsPlayerTurn = !g.IsPlayerTurn
	}
	vector.DrawFilledCircle(screen, g.Ball.X, g.Ball.Y, g.Ball.Radius, color.RGBA{255, 255, 255, 255}, false)
}

func (g *Game) HandlePlayer(screen *ebiten.Image) {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.Player.Y -= 5
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.Player.Y += 5
	}
	vector.DrawFilledRect(screen, g.Player.X, g.Player.Y, g.Player.Width, g.Player.Height, color.RGBA{255, 255, 255, 255}, true)
}

func (g *Game) HandleBot(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 20, 50, 6, 30, color.RGBA{255, 255, 255, 255}, true)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.HandleBot(screen)
	g.HandleBall(screen)
	g.HandlePlayer(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func (g *Game) Init() {
	g.Ball.X = float32(SCREEN_WIDTH) / 2
	g.Ball.Y = float32(SCREEN_HEIGHT) / 2
	g.Ball.Radius = 3
	g.BallSpeed = 2.5
	g.IsPlayerTurn = true

	g.Player.X = 294
	g.Player.Y = 50
	g.Player.Width = 6
	g.Player.Height = 30
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
