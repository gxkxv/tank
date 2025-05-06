package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
)

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}
func main() {
	g := &Game{}

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Tank Game")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
