package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Explode struct {
	x, y     float32
	step     int
	diameter []int
	live     bool
	game     *Game
}

const (
	EXPLOSION_STEPS = 10
)

func NewExplode(x, y float32, g *Game) *Explode {
	return &Explode{
		x:        x,
		y:        y,
		step:     0,
		diameter: []int{4, 7, 12, 18, 26, 32, 49, 30, 14, 6},
		live:     true,
		game:     g,
	}
}
func (e *Explode) Draw(screen *ebiten.Image) {
	if !e.live {
		return
	}

	if e.step >= len(e.diameter) {
		e.live = false
		return
	}

	explosionColor := color.RGBA{255, 165, 0, 255} // Orange color for explosion
	vector.DrawFilledCircle(screen, e.x, e.y, float32(e.diameter[e.step]), explosionColor, false)

	e.step++
}

func (e *Explode) IsLive() bool {
	return e.live
}
