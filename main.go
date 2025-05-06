package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"time"
)

const GAME_WIDTH = 800
const GAME_HEIGHT = 600

type Game struct {
	tank     *Tank
	missiles []*Missile
	lastShot time.Time
}

func NewTank(x, y float32) *Tank {
	return &Tank{
		x:         x,
		y:         y,
		speed:     5,
		direction: STOP,
	}
}

func NewGame() *Game {
	return &Game{tank: NewTank(GAME_WIDTH/2, GAME_HEIGHT/2)}
}

func NewMissile(x, y float32, direction Direction) *Missile {
	return &Missile{
		x:         x,
		y:         y,
		direction: direction,
	}
}

func (g *Game) Update() error {
	g.tank.Update()

	if ebiten.IsKeyPressed(ebiten.KeyW) && time.Since(g.lastShot) > time.Millisecond*100 {
		g.lastShot = time.Now()
		g.missiles = append(g.missiles, NewMissile(g.tank.x, g.tank.y, g.tank.direction))
	}
	for i := len(g.missiles) - 1; i >= 0; i-- {
		g.missiles[i].Move()
		if g.missiles[i].x < 0 || g.missiles[i].x > GAME_WIDTH || g.missiles[i].y < 0 || g.missiles[i].y > GAME_HEIGHT {
			g.missiles = append(g.missiles[:i], g.missiles[i+1:]...)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Use arrow keys to navigate.")
	g.tank.Draw(screen)
	for _, missile := range g.missiles {
		missile.Draw(screen)
	}
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return GAME_WIDTH, GAME_HEIGHT
}
func main() {
	g := NewGame()
	// set window size
	ebiten.SetWindowSize(GAME_WIDTH, GAME_HEIGHT)
	ebiten.SetWindowTitle("Tank Game")

	// start the game
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
