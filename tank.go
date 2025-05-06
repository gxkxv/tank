package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Game struct {
	x, y      float32 // our position
	offScreen *ebiten.Image
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.y -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.y += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.x -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.x += 5
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.offScreen.Clear()
	g.offScreen.Fill(color.RGBA{255, 255, 255, 255})
	tankColor := color.RGBA{255, 0, 0, 255}
	vector.DrawFilledCircle(g.offScreen, g.x, g.y, 15, tankColor, false) //draw a circle

	screen.DrawImage(g.offScreen, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600 // screen size
}
