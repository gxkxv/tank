package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

// Создаем структуру Rectangle для обработки столкновений
type Rectangle struct {
	x, y, width, height float32
}

// Метод для проверки пересечения двух прямоугольников
func (r *Rectangle) Intersects(other *Rectangle) bool {
	return r.x < other.x+other.width &&
		r.x+r.width > other.x &&
		r.y < other.y+other.height &&
		r.y+r.height > other.y
}

type Missile struct {
	x, y      float32
	direction Direction
	active    bool
	game      Game
}

const (
	MissileRadius = 6
	MissileSpeed  = 10
)

func (m *Missile) Draw(screen *ebiten.Image) {
	if !m.active {
		return
	}
	missileColor := color.RGBA{255, 255, 255, 255}
	vector.DrawFilledCircle(screen, m.x, m.y, MissileRadius, missileColor, false)
}

// Move method for bullet`s move
func (m *Missile) Move() {
	if !m.active {
		return
	}
	switch m.direction {
	case LEFT:
		m.x -= MissileSpeed
	case LEFT_UP:
		m.x -= MissileSpeed
		m.y -= MissileSpeed
	case UP:
		m.y -= MissileSpeed
	case RIGHT_UP:
		m.x += MissileSpeed
		m.y -= MissileSpeed
	case RIGHT:
		m.x += MissileSpeed
	case RIGHT_DOWN:
		m.x += MissileSpeed
		m.y += MissileSpeed
	case DOWN:
		m.y += MissileSpeed
	case LEFT_DOWN:
		m.x -= MissileSpeed
		m.y += MissileSpeed
	}
	if m.x < 0 || m.x > GAME_WIDTH || m.y < 0 || m.y > GAME_HEIGHT {
		m.active = false
	}
}

// GetRect возвращает прямоугольник для снаряда
func (m *Missile) GetRect() *Rectangle {
	return &Rectangle{x: m.x, y: m.y, width: MissileRadius * 2, height: MissileRadius * 2}
}

// hitTank проверяет, столкнулся ли снаряд с танком
func (m *Missile) hitTank(t *Tank) bool {
	if m.GetRect().Intersects(t.GetRect()) && t.IsLive() {
		t.SetLive(false) // Уничтожение танка
		m.active = false // Уничтожение снаряда
		m.game.AddExplosion(t.x, t.y)
		return true
	}
	return false
}

func (m *Missile) hitTanks(tanks []*Tank) ([]*Tank, bool) {
	for i := 0; i < len(tanks); i++ {
		if m.hitTank(tanks[i]) {
			// Удаляем танк из среза
			tanks = append(tanks[:i], tanks[i+1:]...)
			fmt.Println("Enemies left:", len(tanks)) // Проверка количества врагов
			return tanks, true
		}
	}
	return tanks, false
}

//func (m *Missile) hitTanks(tanks []*Tank) {
//	for i := len(tanks) - 1; i >= 0; i-- {
//		if m.getRect().Intersects(tanks[i].getRect()) && tanks[i].isLive() {
//			tanks[i].setLive(false)
//			m.setActive(false) // Уничтожаем снаряд
//			explode := NewExplode(tanks[i].x, tanks[i].y, m.tc)
//			m.tc.explodes = append(m.tc.explodes, explode)
//
//			// Удаляем уничтоженный танк
//			tanks = append(tanks[:i], tanks[i+1:]...)
//			break
//		}
//	}
//}
