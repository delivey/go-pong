package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) Update() error {
	g.HandleBot()
	g.HandleBall()
	g.HandlePlayer()
	return nil
}

func (g *Game) HandleBallCollisions() {
	ball := g.Ball
	player := g.Player

	collisionX := ball.X+ball.Width >= player.X &&
		player.X+player.Width >= ball.X

	collisionY := ball.Y+ball.Height >= player.Y &&
		player.Y+player.Height >= ball.Y
	if collisionX && collisionY {
		g.BallSpeed *= -1
	}
}

func (g *Game) HandleBall() {
	g.Ball.X += g.BallSpeed
	g.HandleBallCollisions()
	// Handle point ending
	if g.Ball.X > float32(SCREEN_WIDTH) {
		g.Player.Score++
		g.IsPlayerTurn = !g.IsPlayerTurn
		g.BallSpeed *= -1
		g.Ball.X = float32(SCREEN_WIDTH) / 2
	}
	if g.Ball.X < 0 {
		g.Bot.Score++
		g.IsPlayerTurn = !g.IsPlayerTurn
		g.BallSpeed *= -1
		g.Ball.X = float32(SCREEN_WIDTH) / 2
	}
}

func (g *Game) HandlePlayer() {
	if g.Player.IsMovingUp {
		g.Player.Y -= g.Player.Speed
	}
	if g.Player.IsMovingDown {
		g.Player.Y += g.Player.Speed
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.Player.IsMovingUp = true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.Player.IsMovingDown = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowUp) {
		g.Player.IsMovingUp = false
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowDown) {
		g.Player.IsMovingDown = false
	}
}

func (g *Game) HandleBot() {
}

func (g *Game) Draw(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprint(g.Player.Score), g.Font, 106, 36, color.White)
	text.Draw(screen, fmt.Sprint(g.Bot.Score), g.Font, 208, 36, color.White)

	vector.DrawFilledRect(screen, 20, 50, 6, 30, color.RGBA{255, 255, 255, 255}, true)
	vector.DrawFilledRect(screen, g.Ball.X, g.Ball.Y, g.Ball.Width, g.Ball.Height, color.RGBA{255, 255, 255, 255}, true)
	vector.DrawFilledRect(screen, g.Player.X, g.Player.Y, g.Player.Width, g.Player.Height, color.RGBA{255, 255, 255, 255}, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func (g *Game) Init() {
	g.Ball.X = float32(SCREEN_WIDTH) / 2
	g.Ball.Y = float32(SCREEN_HEIGHT) / 2
	g.Ball.Width = 4.5
	g.Ball.Height = 4.5
	g.BallSpeed = 2.5
	g.IsPlayerTurn = true

	g.Player.X = 294
	g.Player.Y = 50
	g.Player.Width = 6
	g.Player.Height = 30
	g.Player.Speed = 3.5

	g.Font = GetFont()
}

func main() {
	ebiten.SetWindowSize(960, 720)
	ebiten.SetWindowTitle("Pong")
	game := &Game{}
	game.Init()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
