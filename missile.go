package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Missile struct {
	x, y      float32
	direction Direction
	active    bool
}

const (
	MissileRadius = 6
	MissileSpeed  = 10
)

func (m *Missile) Draw(screen *ebiten.Image) {
	if !m.active {
		return
	}
	missileColor := color.RGBA{255, 255, 255, 255}
	vector.DrawFilledCircle(screen, m.x, m.y, MissileRadius, missileColor, false)
}

// Move method for bullet`s move
func (m *Missile) Move() {
	if !m.active {
		return
	}
	switch m.direction {
	case LEFT:
		m.x -= MissileSpeed
	case LEFT_UP:
		m.x -= MissileSpeed
		m.y -= MissileSpeed
	case UP:
		m.y -= MissileSpeed
	case RIGHT_UP:
		m.x += MissileSpeed
		m.y -= MissileSpeed
	case RIGHT:
		m.x += MissileSpeed
	case RIGHT_DOWN:
		m.x += MissileSpeed
		m.y += MissileSpeed
	case DOWN:
		m.y += MissileSpeed
	case LEFT_DOWN:
		m.x -= MissileSpeed
		m.y += MissileSpeed
	}
	if m.x < 0 || m.x > GAME_WIDTH || m.y < 0 || m.y > GAME_HEIGHT {
		m.active = false
	}
}
