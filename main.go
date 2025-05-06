package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const GAME_WIDTH = 800
const GAME_HEIGHT = 600

func main() {
	g := &Game{
		offScreen: ebiten.NewImage(640, 480),
	}

	// set window size
	ebiten.SetWindowSize(GAME_WIDTH, GAME_HEIGHT)
	ebiten.SetWindowTitle("Tank Game")

	// start the game
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
