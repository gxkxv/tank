package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"time"
)

type Direction int
type Tank struct {
	x, y          float32 // our position
	speed         float32 // speed
	direction     Direction
	ptDir         Direction
	fireRequested bool
	lastFiredAt   time.Time
	prevFireKey   bool // last status W
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
	fireKeyPressed := ebiten.IsKeyPressed(ebiten.KeyW)
	now := time.Now()
	// make some delay
	if fireKeyPressed && !t.prevFireKey && now.Sub(t.lastFiredAt) > 800*time.Millisecond {
		t.fireRequested = true
		t.lastFiredAt = now
	}
	//cant hold
	t.prevFireKey = fireKeyPressed
	if t.direction != STOP {
		t.ptDir = t.direction
	}
	t.move()
}

func (t *Tank) move() {
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
	}

	if t.direction != STOP {
		t.ptDir = t.direction
	}
}

func (t *Tank) Draw(screen *ebiten.Image) {
	tankColor := color.RGBA{255, 0, 0, 255}
	vector.DrawFilledCircle(screen, t.x, t.y, 15, tankColor, false) //draw a circle

	cx, cy := t.x, t.y
	var ex, ey float32

	switch t.ptDir {
	case LEFT:
		ex, ey = cx-25, cy
	case LEFT_UP:
		ex, ey = cx-20, cy-20
	case UP:
		ex, ey = cx, cy-25
	case RIGHT_UP:
		ex, ey = cx+20, cy-20
	case RIGHT:
		ex, ey = cx+25, cy
	case RIGHT_DOWN:
		ex, ey = cx+20, cy+20
	case DOWN:
		ex, ey = cx, cy+25
	case LEFT_DOWN:
		ex, ey = cx-20, cy+20
	default:
		ex, ey = cx, cy
	}

	vector.StrokeLine(screen, cx, cy, ex, ey, 3, tankColor, false)
}

func (t *Tank) Fire() *Missile {
	t.fireRequested = false
	if t.direction == STOP {
		return nil
	}
	mx := t.x + 15/2 - MissileRadius/2
	my := t.y + 15/2 - MissileRadius/2
	return NewMissile(mx, my, t.direction)
}
