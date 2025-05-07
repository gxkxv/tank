package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math/rand"
	"time"
)

type Direction int
type Tank struct {
	x, y          float32 // our position
	speed         float32 // speed
	direction     Direction
	ptDir         Direction
	fireRequested bool
	lastFiredAt   time.Time
	prevFireKey   bool // last status W
	live          bool // Жив ли танк
	good          bool
	step          int
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

const (
	STOP = iota
	LEFT
	LEFT_UP
	UP
	RIGHT_UP
	RIGHT
	RIGHT_DOWN
	DOWN
	LEFT_DOWN
)

func (t *Tank) Update() {
	if !t.good { // Для вражеских танков
		t.randomMove() // Двигаемся в случайном направлении
	}
	fireKeyPressed := false // Делаем выстрел в случайный момент времени для врага

	if t.good {
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			if ebiten.IsKeyPressed(ebiten.KeyUp) {
				t.direction = LEFT_UP
			} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
				t.direction = LEFT_DOWN
			} else {
				t.direction = LEFT
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			if ebiten.IsKeyPressed(ebiten.KeyUp) {
				t.direction = RIGHT_UP
			} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
				t.direction = RIGHT_DOWN
			} else {
				t.direction = RIGHT
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
			t.direction = UP
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			t.direction = DOWN
		} else {
			t.direction = STOP
		}
		fireKeyPressed = ebiten.IsKeyPressed(ebiten.KeyW)
	}

	now := time.Now()
	// make some delay
	if fireKeyPressed && !t.prevFireKey && now.Sub(t.lastFiredAt) > 1*time.Millisecond {
		t.fireRequested = true
		t.lastFiredAt = now
	}
	//cant hold
	t.prevFireKey = fireKeyPressed
	if t.direction != STOP {
		t.ptDir = t.direction
	}
	t.move()
}
func (t *Tank) randomMove() {
	if t.step == 0 {
		t.step = r.Intn(12) + 3 // Генерируем случайное количество шагов
		dirIdx := r.Intn(9)     // Случайное направление
		t.direction = Direction(dirIdx)
	}

	t.step--
	if r.Intn(40) > 38 { // С вероятностью 5% танк стрельнет
		t.fireRequested = true
	}
	if t.x < 0 {
		t.x = 0
	}
	if t.y < 0 {
		t.y = 0
	}
	if t.x > GAME_WIDTH-15 { // Минус радиус танка (15 пикселей)
		t.x = GAME_WIDTH - 15
	}
	if t.y > GAME_HEIGHT-15 { // Минус радиус танка (15 пикселей)
		t.y = GAME_HEIGHT - 15
	}

}
func (t *Tank) move() {
	switch t.direction {
	case LEFT:
		t.x -= t.speed
	case LEFT_UP:
		t.x -= t.speed
		t.y -= t.speed
	case UP:
		t.y -= t.speed
	case RIGHT_UP:
		t.x += t.speed
		t.y -= t.speed
	case RIGHT:
		t.x += t.speed
	case RIGHT_DOWN:
		t.x += t.speed
		t.y += t.speed
	case DOWN:
		t.y += t.speed
	case LEFT_DOWN:
		t.x -= t.speed
		t.y += t.speed
	}

	if t.direction != STOP {
		t.ptDir = t.direction
	}
	if t.x < 0 {
		t.x = 0
	}
	if t.y < 0 {
		t.y = 0
	}
	if t.x > GAME_WIDTH-15 { // Минус радиус танка (15 пикселей)
		t.x = GAME_WIDTH - 15
	}
	if t.y > GAME_HEIGHT-15 { // Минус радиус танка (15 пикселей)
		t.y = GAME_HEIGHT - 15
	}

}

func (t *Tank) Draw(screen *ebiten.Image) {
	if !t.live {
		return
	}

	tankColor := color.RGBA{255, 0, 0, 255}

	if !t.good {
		tankColor = color.RGBA{0, 255, 0, 255}
	}
	vector.DrawFilledCircle(screen, t.x, t.y, 15, tankColor, false) //draw a circle

	cx, cy := t.x, t.y
	var ex, ey float32

	switch t.ptDir {
	case LEFT:
		ex, ey = cx-25, cy
	case LEFT_UP:
		ex, ey = cx-20, cy-20
	case UP:
		ex, ey = cx, cy-25
	case RIGHT_UP:
		ex, ey = cx+20, cy-20
	case RIGHT:
		ex, ey = cx+25, cy
	case RIGHT_DOWN:
		ex, ey = cx+20, cy+20
	case DOWN:
		ex, ey = cx, cy+25
	case LEFT_DOWN:
		ex, ey = cx-20, cy+20
	default:
		ex, ey = cx, cy
	}

	vector.StrokeLine(screen, cx, cy, ex, ey, 3, tankColor, false)
}
func (t *Tank) Fire() *Missile {
	t.fireRequested = false
	if t.direction == STOP {
		return nil
	}
	mx := t.x + 15/2 - MissileRadius/2
	my := t.y + 15/2 - MissileRadius/2
	return &Missile{
		x:         mx,
		y:         my,
		direction: t.ptDir,
		active:    true,
		good:      t.good,
	}
}

func (t *Tank) SetLive(live bool) {
	t.live = live
}

func (t *Tank) IsLive() bool {
	return t.live
}

// GetRect возвращает прямоугольник, представляющий границы танка
func (t *Tank) GetRect() *Rectangle {
	// Возвращаем прямоугольник с границами танка
	return &Rectangle{x: t.x, y: t.y, width: 15, height: 15}
}

//func (t *Tank) UpdateAutomatically() {
//	// Или случайное движение (например, вправо или вниз)
//	if rand.Float32() < 0.25 {
//		t.x += t.speed
//	} else if rand.Float32() > 0.25 && rand.Float32() < 0.5 {
//		t.y += t.speed
//	} else if rand.Float32() > 0.5 && rand.Float32() < 0.75 {
//		t.y -= t.speed
//	} else if rand.Float32() > 0.75 && rand.Float32() < 1 {
//		t.x -= t.speed
//	}
//}
