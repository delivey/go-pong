package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) Update() error {
	g.HandleBot()
	g.HandleBall()
	g.HandlePlayer()
	g.CurrentFPS = ebiten.ActualFPS()
	return nil
}

func CheckPaddleCollision(paddle Paddle, ball Ball) bool {
	collisionX := ball.X+ball.Width >= paddle.X &&
		paddle.X+paddle.Width >= ball.X

	collisionY := ball.Y+ball.Height >= paddle.Y &&
		paddle.Y+paddle.Height >= ball.Y
	return collisionX && collisionY
}

func GuessBallYPosition(ball Ball, endX float32) float32 {
	currentX := ball.X
	currentY := ball.Y
	for {
		currentX += ball.SpeedX
		currentY += ball.SpeedY
		fmt.Printf("CurrentX: %.2f\n", currentX)
		if currentX >= endX {
			break
		}
	}
	return currentY
}

func GetYMultiplier(paddle Paddle, ball Ball) float32 {
	var result float32
	if ball.Y > paddle.Y {
		result = (paddle.Y + paddle.Height/2) / 100
	} else {
		result = (paddle.Y - paddle.Height/2) / -100
	}
	if result > 1 {
		return 1
	}
	if result < -1 {
		return -1
	}
	return result
}

func (g *Game) HandleBallCollisions() {
	botCollision := CheckPaddleCollision(g.Bot.Paddle, g.Ball)
	playerCollision := CheckPaddleCollision(g.Player.Paddle, g.Ball)
	var collidedPaddle Paddle
	if playerCollision {
		collidedPaddle = g.Player.Paddle
	}
	if botCollision {
		collidedPaddle = g.Bot.Paddle
	}
	if botCollision || playerCollision {
		g.Ball.SpeedX *= -1
		yMultiplier := GetYMultiplier(collidedPaddle, g.Ball)
		g.Ball.SpeedY += yMultiplier
	}
}

func (g *Game) ResetBall() {
	g.IsPlayerTurn = !g.IsPlayerTurn
	g.Ball.SpeedX *= -1
	g.Ball.SpeedY = 0
	g.Ball.Y = float32(SCREEN_HEIGHT) / 2
	g.Ball.X = float32(SCREEN_WIDTH) / 2
}

func (g *Game) HandleBall() {
	g.Ball.X += g.Ball.SpeedX
	g.Ball.Y += g.Ball.SpeedY
	g.HandleBallCollisions()
	// Handle point win
	if g.Ball.X > float32(SCREEN_WIDTH) {
		g.Player.Score++
		g.ResetBall()
	}
	if g.Ball.X < 0 {
		g.Bot.Score++
		g.ResetBall()
	}

	// Handle ball bounce
	if g.Ball.Y+g.Ball.Height > float32(SCREEN_HEIGHT) || g.Ball.Y-g.Ball.Height < 0 {
		g.Ball.SpeedY *= -1
	}
}

func (g *Game) HandlePlayer() {

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
	if g.Player.Y >= float32(SCREEN_HEIGHT)-g.Player.Height && g.Player.IsMovingDown {
		return
	}
	if g.Player.Y <= 0 && g.Player.IsMovingUp {
		return
	}
	if g.Player.IsMovingUp {
		g.Player.Y -= g.Player.Speed
	}
	if g.Player.IsMovingDown {
		g.Player.Y += g.Player.Speed
	}
}

func (g *Game) HandleBot() {
	if g.Ball.X < g.Bot.X {
		return
	}
	if g.Ball.SpeedX > 0 {
		return
	}

	// Predict where they ball will be and move there
	predictedY := g.Ball.Y
	if math.Abs(float64(predictedY-g.Bot.Y)) <= float64(g.Bot.Speed) {
		return
	}
	if predictedY > g.Bot.Y {
		g.Bot.Y += g.Bot.Speed
		if g.Bot.Y >= float32(SCREEN_HEIGHT)-g.Bot.Height {
			g.Bot.Y -= g.Bot.Speed
		}
	} else if predictedY < g.Bot.Y {
		g.Bot.Y -= g.Bot.Speed
		if g.Bot.Y <= 0 {
			g.Bot.Y += g.Bot.Speed
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprint(g.Player.Score), g.Font, 106, 36, color.White)
	text.Draw(screen, fmt.Sprint(g.Bot.Score), g.Font, 208, 36, color.White)
	fpsText := fmt.Sprintf("FPS: %.2f", g.CurrentFPS)
	text.Draw(screen, fpsText, g.Font, 250, 20, color.White)

	vector.DrawFilledRect(screen, g.Bot.X, g.Bot.Y, g.Bot.Width, g.Bot.Height, color.RGBA{255, 255, 255, 255}, true)
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
	g.Ball.SpeedX = 4
	g.Ball.SpeedY = 0
	g.IsPlayerTurn = true

	g.Player.X = 294
	g.Player.Y = float32(SCREEN_HEIGHT) / 2
	g.Player.Width = 6
	g.Player.Height = 30
	g.Player.Speed = 3.5

	g.Bot.X = 26
	g.Bot.Y = float32(SCREEN_HEIGHT) / 2
	g.Bot.Width = 6
	g.Bot.Height = 30
	g.Bot.Speed = 3.5

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
