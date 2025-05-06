package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Direction int
type Tank struct {
	x, y      float32 // our position
	speed     float32 // speed
	direction Direction
}

const (
	STOP = iota
	LEFT
	LEFT_UP
	UP
	RIGHT_UP
	RIGHT
	RIGHT_DOWN
	DOWN
	LEFT_DOWN
)

func (t *Tank) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			t.direction = LEFT_UP
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			t.direction = LEFT_DOWN
		} else {
			t.direction = LEFT
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			t.direction = RIGHT_UP
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			t.direction = RIGHT_DOWN
		} else {
			t.direction = RIGHT
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		t.direction = UP
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		t.direction = DOWN
	} else {
		t.direction = STOP
	}

	switch t.direction {
	case LEFT:
		t.x -= t.speed
	case LEFT_UP:
		t.x -= t.speed
		t.y -= t.speed
	case UP:
		t.y -= t.speed
	case RIGHT_UP:
		t.x += t.speed
		t.y -= t.speed
	case RIGHT:
		t.x += t.speed
	case RIGHT_DOWN:
		t.x += t.speed
		t.y += t.speed
	case DOWN:
		t.y += t.speed
	case LEFT_DOWN:
		t.x -= t.speed
		t.y += t.speed
	case STOP:
	}
}

func (t *Tank) Draw(screen *ebiten.Image) {
	tankColor := color.RGBA{255, 0, 0, 255}
	vector.DrawFilledCircle(screen, t.x, t.y, 15, tankColor, false) //draw a circle
}
