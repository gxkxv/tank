package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Missile struct {
	x, y      float32
	direction Direction
}

const (
	XSPEED = 10
	YSPEED = 10
)

func (m *Missile) Draw(screen *ebiten.Image) {
	missileColor := color.RGBA{255, 255, 255, 255}
	vector.DrawFilledCircle(screen, m.x, m.y, 5, missileColor, false)
	m.Move()
}

// Move method for bullet`s move
func (m *Missile) Move() {
	switch m.direction {
	case LEFT:
		m.x -= XSPEED
	case LEFT_UP:
		m.x -= XSPEED
		m.y -= YSPEED
	case UP:
		m.y -= YSPEED
	case RIGHT_UP:
		m.x += XSPEED
		m.y -= YSPEED
	case RIGHT:
		m.x += XSPEED
	case RIGHT_DOWN:
		m.x += XSPEED
		m.y += YSPEED
	case DOWN:
		m.y += YSPEED
	case LEFT_DOWN:
		m.x -= XSPEED
		m.y += YSPEED
	default:
		m.x += XSPEED
	}

}
