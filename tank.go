package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
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
	game          *Game
	lastSuperTime time.Time
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
	if ebiten.IsKeyPressed(ebiten.KeyA) && time.Since(t.lastSuperTime) > 5*time.Second {
		t.lastSuperTime = time.Now()
		t.PerformSuperAttack(t.game) // передай ссылку на Game
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
func (t *Tank) PerformSuperAttack(g *Game) {
	if !t.good || !t.live {
		return
	}

	const (
		count     = 5           // количество взрывов в волне
		spacing   = float32(40) // расстояние между взрывами
		radiusHit = float32(30) // радиус поражения
	)

	// Начальная позиция — центр танка
	startX, startY := t.x, t.y
	dirX, dirY := directionVector(t.ptDir)

	for i := 1; i <= count; i++ {
		ex := startX + dirX*spacing*float32(i)
		ey := startY + dirY*spacing*float32(i)

		g.explodes = append(g.explodes, NewExplosion(ex, ey))

		// Уничтожаем врагов в радиусе
		var aliveEnemies []*Tank
		for _, e := range g.enemyTanks {
			if distance(ex, ey, e.x, e.y) < radiusHit {
				e.live = false
				g.explodes = append(g.explodes, NewExplosion(e.x, e.y))
			} else {
				aliveEnemies = append(aliveEnemies, e)
			}
		}
		g.enemyTanks = aliveEnemies
	}
}

func distance(x1, y1, x2, y2 float32) float32 {
	dx := x2 - x1
	dy := y2 - y1
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}
func directionVector(dir Direction) (float32, float32) {
	switch dir {
	case LEFT:
		return -1, 0
	case LEFT_UP:
		return -0.7, -0.7
	case UP:
		return 0, -1
	case RIGHT_UP:
		return 0.7, -0.7
	case RIGHT:
		return 1, 0
	case RIGHT_DOWN:
		return 0.7, 0.7
	case DOWN:
		return 0, 1
	case LEFT_DOWN:
		return -0.7, 0.7
	default:
		return 0, -1 // по умолчанию вверх
	}
}
func drawSuperAttackIcon(screen *ebiten.Image, t *Tank) {
	x, y := GAME_WIDTH-60, GAME_HEIGHT-60

	// Фон и рамка
	ebitenutil.DrawRect(screen, float64(x), float64(y), 50, 50, color.RGBA{80, 80, 80, 255})

	// Прозрачная заливка — оставшееся время кулдауна
	cd := time.Since(t.lastSuperTime)
	cooldown := 5 * time.Second

	if cd < cooldown {
		// вычисляем сколько залить (чем меньше осталось — тем выше заливка)
		h := 50 * (1 - float64(cd)/float64(cooldown))
		ebitenutil.DrawRect(screen, float64(x), float64(y)+h, 50, 50-h, color.RGBA{255, 0, 0, 180})
	} else {
		// готов к атаке — подсветим зелёным
		ebitenutil.DrawRect(screen, float64(x), float64(y), 50, 50, color.RGBA{0, 255, 0, 80})
	}
}
