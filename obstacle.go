package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

type ObstacleType int

const (
	Wall ObstacleType = iota
	Rock
	Bush
	Building
)

type Obstacle struct {
	X, Y, Width, Height float64
	Type                ObstacleType
}

func NewObstacle(x, y, w, h float64, obsType ObstacleType) *Obstacle {
	return &Obstacle{
		X:      x,
		Y:      y,
		Width:  w,
		Height: h,
		Type:   obsType,
	}
}

func (o *Obstacle) Draw(screen *ebiten.Image) {
	switch o.Type {
	case Wall:
		// Кирпичная стена
		ebitenutil.DrawRect(screen, o.X, o.Y, o.Width, o.Height, color.RGBA{255, 255, 255, 255})

	case Rock:
		// Серый камень с неровными краями
		ebitenutil.DrawCircle(screen, o.X+o.Width/2, o.Y+o.Height/2, o.Width/2, color.RGBA{100, 100, 100, 255})
	case Bush:
		// Зеленый куст
		ebitenutil.DrawCircle(screen, o.X+o.Width/2, o.Y+o.Height/2, o.Width/2, color.RGBA{0, 100, 0, 200})
	case Building:
		// Серое здание с окнами
		ebitenutil.DrawRect(screen, o.X, o.Y, o.Width, o.Height, color.RGBA{70, 70, 70, 255})
		drawWindows(screen, o.X, o.Y, o.Width, o.Height)
	}
}

func drawWindows(screen *ebiten.Image, x, y, w, h float64) {
	windowW, windowH := 10.0, 15.0
	margin := 5.0
	for i := margin; i < w-margin; i += windowW + margin {
		for j := margin; j < h-margin; j += windowH + margin {
			ebitenutil.DrawRect(
				screen,
				x+i,
				y+j,
				windowW,
				windowH,
				color.RGBA{200, 200, 100, 255},
			)
		}
	}
}

func (o *Obstacle) Bounds() (x, y, w, h float64) {
	return o.X, o.Y, o.Width, o.Height
}
