package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Tank Game!") // shows the text
	tankColor := color.RGBA{255, 0, 0, 255}
	vector.DrawFilledCircle(screen, 50, 50, 30, tankColor, false) //draw a circle

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600 // screen size
}
func main() {
	g := &Game{}

	// set window size
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Tank Game")

	// start the game
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
