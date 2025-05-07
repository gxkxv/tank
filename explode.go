package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

type Explode struct {
	x, y      float32
	radius    float32
	maxradius float32
	duration  int
}

func NewExplosion(x, y float32) *Explode {
	return &Explode{
		x:         x,
		y:         y,
		maxradius: 30,
		duration:  20,
	}
}

func (e *Explode) Update() bool {
	e.radius = e.maxradius * (1 - float32(e.duration)/20)
	e.duration--
	return e.duration > 0
}

func (e *Explode) Draw(screen *ebiten.Image) {
	ebitenutil.DrawCircle(screen, float64(e.x), float64(e.y), float64(e.radius), color.RGBA{255, 150, 0, 150})
}
