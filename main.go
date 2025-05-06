package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

const GAME_WIDTH = 800
const GAME_HEIGHT = 600

type Game struct {
	tank      *Tank
	enemyTank *Tank
	missiles  []*Missile
}

func NewTank(x, y float32) *Tank {
	return &Tank{
		x:         x,
		y:         y,
		speed:     5,
		direction: STOP,
		ptDir:     DOWN,
	}
}

func NewGame() *Game {
	return &Game{
		tank:      NewTank(GAME_WIDTH/2, GAME_HEIGHT/2), // Игровой танк
		enemyTank: NewTank(GAME_WIDTH/4, GAME_HEIGHT/4), // Вражеский танк (фиксированная позиция для примера)
	}
}

func NewMissile(x, y float32, direction Direction) *Missile {
	return &Missile{
		x:         x,
		y:         y,
		direction: direction,
		active:    true,
	}
}

func (g *Game) Update() error {
	// Обновляем наш танк
	g.tank.Update()

	// Обновляем врага
	g.enemyTank.Update()

	// Выстрелы нашего танка
	if g.tank.fireRequested {
		if m := g.tank.Fire(); m != nil {
			g.missiles = append(g.missiles, m)
		}
	}

	// Выстрелы врага
	if g.enemyTank.fireRequested {
		if m := g.enemyTank.Fire(); m != nil {
			g.missiles = append(g.missiles, m)
		}
	}

	// Двигаем все снаряды
	for _, m := range g.missiles {
		m.Move()
	}

	// Очищаем неактивные снаряды
	var activeMissiles []*Missile
	for _, m := range g.missiles {
		if m.active {
			activeMissiles = append(activeMissiles, m)
		}
	}
	g.missiles = activeMissiles

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Use arrow keys to navigate.")
	g.tank.Draw(screen)
	g.enemyTank.Draw(screen)

	// Рисуем все снаряды
	for _, missile := range g.missiles {
		missile.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return GAME_WIDTH, GAME_HEIGHT
}

func main() {
	// Запускаем игру
	g := NewGame()
	ebiten.SetWindowSize(GAME_WIDTH, GAME_HEIGHT)
	ebiten.SetWindowTitle("Tank Game")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
