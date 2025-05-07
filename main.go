package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math/rand"
	"strconv"
)

const GAME_WIDTH = 800
const GAME_HEIGHT = 600

type Game struct {
	tank       *Tank
	enemyTanks []*Tank
	missiles   []*Missile
	explodes   []*Explode
}

func NewTank(x, y float32) *Tank {
	return &Tank{
		x:         x,
		y:         y,
		speed:     5,
		direction: STOP,
		ptDir:     DOWN,
		live:      true,
	}
}

func NewGame() *Game {
	return &Game{
		tank:       NewTank(GAME_WIDTH/2, GAME_HEIGHT/2), //Игровой танк
		enemyTanks: []*Tank{},                            // Вражеский танк

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

func (g *Game) AddExplosion(x, y float32) {
	explode := NewExplode(x, y, g)           // Создаем взрыв
	g.explodes = append(g.explodes, explode) // Добавляем в список взрывов
}

func (g *Game) Update() error {
	// Обновляем наш танк
	g.tank.Update()

	// Обновляем врага
	//g.enemyTank.UpdateAutomatically()

	// Выстрелы нашего танка
	if g.tank.fireRequested {
		if m := g.tank.Fire(); m != nil {
			g.missiles = append(g.missiles, m)
		}
	}

	// Выстрелы врага

	for _, enemy := range g.enemyTanks {
		if enemy.fireRequested {
			if m := enemy.Fire(); m != nil {
				g.missiles = append(g.missiles, m)
			}
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

	var activeExplosions []*Explode
	for _, e := range g.explodes {
		if e.IsLive() {
			activeExplosions = append(activeExplosions, e)
		}
	}
	g.explodes = activeExplosions
	for _, m := range g.missiles {
		var hit bool
		g.enemyTanks, hit = m.hitTanks(g.enemyTanks) // Получаем обновленный срез врагов
		if hit {
			// Можно добавить дополнительные действия при попадании снаряда в танк
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Use arrow keys to navigate.")
	g.tank.Draw(screen)

	for _, enemyTank := range g.enemyTanks {
		enemyTank.Draw(screen)
	}
	// Рисуем все снаряды
	for _, missile := range g.missiles {
		missile.Draw(screen)
	}
	for _, explode := range g.explodes {
		explode.Draw(screen)
	}

	// Рисуем количество оставшихся танков
	ebitenutil.DebugPrintAt(screen, "Tanks left: "+strconv.Itoa(len(g.enemyTanks)), GAME_WIDTH/2, GAME_HEIGHT/2)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return GAME_WIDTH, GAME_HEIGHT
}

func (g *Game) launchFrame() {
	// Добавляем 10 врагов
	for i := 0; i < 10; i++ {
		g.enemyTanks = append(g.enemyTanks, NewTank(float32(rand.Intn(GAME_WIDTH)), float32(rand.Intn(GAME_HEIGHT))))
	}

}

func main() {
	// Запускаем игру
	g := NewGame()
	g.launchFrame()
	ebiten.SetWindowSize(GAME_WIDTH, GAME_HEIGHT)
	ebiten.SetWindowTitle("Tank Game")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
