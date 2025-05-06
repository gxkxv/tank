package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

const GAME_WIDTH = 800
const GAME_HEIGHT = 600

type Game struct {
	tank *Tank
}

func NewTank(x, y float32) *Tank {
	return &Tank{
		x:         x,
		y:         y,
		speed:     5,
		direction: STOP,
	}
}
func (g *Game) Update() error {
	g.tank.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Use arrow keys to navigate.")
	g.tank.Draw(screen)
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return GAME_WIDTH, GAME_HEIGHT
}
func main() {
	g := &Game{
		tank: NewTank(GAME_WIDTH/2, GAME_HEIGHT/2), // position will be in the middle of screen
	}
	// set window size
	ebiten.SetWindowSize(GAME_WIDTH, GAME_HEIGHT)
	ebiten.SetWindowTitle("Tank Game")

	// start the game
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
