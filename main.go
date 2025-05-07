package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
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
	obstacles  []*Obstacle
	state      *GameState
}

func NewTank(x, y float32, g bool) *Tank {
	return &Tank{
		x:         x,
		y:         y,
		speed:     5,
		direction: STOP,
		ptDir:     DOWN,
		live:      true,
		good:      g,
	}
}

func NewGame() *Game {
	g := &Game{
		state:      &GameState{},
		tank:       NewTank(GAME_WIDTH/2, GAME_HEIGHT/2, true), //Игровой танк
		enemyTanks: []*Tank{},
		obstacles: []*Obstacle{
			// Здания
			NewObstacle(100, 100, 80, 120, Building),
			NewObstacle(600, 150, 100, 150, Building),

			// Камни
			NewObstacle(200, 400, 40, 40, Rock),
			NewObstacle(300, 500, 50, 50, Rock),

			// Кусты (проходимые, но ограничивающие обзор)
			NewObstacle(400, 200, 60, 60, Bush),
			NewObstacle(500, 300, 80, 80, Bush),
		},
	}
	return g

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
	switch {
	case g.state.Is(StateMenu):
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.state.Set(StatePlaying)
		}
	case g.state.Is(StatePlaying):
		// Обновляем наш танк
		g.tank.Update()

		for _, enemy := range g.enemyTanks {
			enemy.Update()
		}
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
		// В Update после снарядов:
		for _, m := range g.missiles {
			if m.hitTanks(g.enemyTanks, g) {
				continue
			}
			if m.hitTank(g.tank, g) {
				g.state.Set(StateGameOver)
			}
		}
		var activeExplodes []*Explode
		for _, e := range g.explodes {
			if e.Update() {
				activeExplodes = append(activeExplodes, e)
			}
		}
		g.explodes = activeExplodes
		// Очищаем неактивные снаряды
		var activeMissiles []*Missile
		for _, m := range g.missiles {
			if m.active {
				activeMissiles = append(activeMissiles, m)
			}
		}
		g.missiles = activeMissiles
		// Обработка попаданий
		for _, m := range g.missiles {
			if m.hitTanks(g.enemyTanks, g) {
				continue
			}
			if m.hitTank(g.tank, g) {
				g.state.Set(StateGameOver)
			}
		}

		// Удаляем мертвых врагов
		var aliveEnemies []*Tank
		for _, e := range g.enemyTanks {
			if e.live {
				aliveEnemies = append(aliveEnemies, e)
			}
		}
		g.enemyTanks = aliveEnemies

		// Победа, если врагов не осталось
		if len(g.enemyTanks) == 0 {
			g.state.Set(StateWin)
		}

	case g.state.Is(StatePaused):
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.state.Set(StatePlaying)
		}

	case g.state.Is(StateGameOver), g.state.Is(StateWin):
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.resetGame()
			g.state.Set(StatePlaying)
		}
	}
	return nil
}

func (g *Game) resetGame() {
	g.tank = NewTank(GAME_WIDTH/2, GAME_WIDTH/2, true)
	g.enemyTanks = []*Tank{}
	g.missiles = nil
	g.explodes = nil
	g.launchFrame() // ← добавлено
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{30, 30, 30, 255})
	ebitenutil.DebugPrint(screen, "Use arrow keys to navigate.")
	g.tank.Draw(screen)
	for _, obs := range g.obstacles {
		obs.Draw(screen)
	}
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

	switch {
	case g.state.Is(StateMenu):
		DrawTextCentered(screen, "TANKS BATTLE", 400, 200, color.White)
		DrawTextCentered(screen, "Press ENTER to start", 400, 300, color.White)

	case g.state.Is(StatePaused):
		ebitenutil.DrawRect(screen, 0, 0, 800, 600, color.RGBA{0, 0, 0, 150})
		DrawTextCentered(screen, "PAUSED", 400, 300, color.White)
		DrawTextCentered(screen, "Press ESC to continue", 400, 350, color.White)

	case g.state.Is(StateGameOver):
		ebitenutil.DrawRect(screen, 0, 0, 800, 600, color.RGBA{0, 0, 0, 200})
		DrawTextCentered(screen, "GAME OVER", 400, 250, color.White)
		DrawTextCentered(screen, "Press ENTER to restart", 400, 300, color.White)

	case g.state.Is(StateWin):
		ebitenutil.DrawRect(screen, 0, 0, 800, 600, color.RGBA{0, 0, 0, 200})
		DrawTextCentered(screen, "VICTORY!", 400, 250, color.White)
		DrawTextCentered(screen, "Press ENTER to play again", 400, 300, color.White)
	}
	// Рисуем количество оставшихся танков
	ebitenutil.DebugPrintAt(screen, "Tanks left: "+strconv.Itoa(len(g.enemyTanks)), GAME_WIDTH/2, GAME_HEIGHT/2)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return GAME_WIDTH, GAME_HEIGHT
}

func (g *Game) launchFrame() {
	for i := 0; i < 2; i++ {
		g.enemyTanks = append(g.enemyTanks, NewTank(float32(rand.Intn(GAME_WIDTH)), float32(rand.Intn(GAME_HEIGHT)), false))
	}

}

func main() {
	// Запускаем игру
	g := NewGame()
	g.launchFrame()
	ebiten.SetWindowSize(GAME_WIDTH, GAME_HEIGHT)
	ebiten.SetWindowTitle("Tank Game")

	if err := InitFonts(); err != nil {
		log.Fatalf("Failed to load fonts: %v", err)
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
