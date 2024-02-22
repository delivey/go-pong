package main

import "golang.org/x/image/font"

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
	Speed  float32
	Score  int
}

type Player struct {
	Paddle
	IsMovingUp   bool
	IsMovingDown bool
}

type Bot struct {
	Paddle
}

type Game struct {
	Ball         Ball
	IsPlayerTurn bool
	BallSpeed    float32
	Player       Player
	Bot          Bot
	Font         font.Face
}
