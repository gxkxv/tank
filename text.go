package main

import (
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"image/color"
	"log"
)

var (
	regularFont font.Face
)

// InitFonts инициализирует шрифты при запуске игры
func InitFonts() error {

	tt, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Printf("Failed to parse font: %v", err)
		return err
	}

	// Создаем font.Face с нужными параметрами
	regularFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,               // Размер шрифта
		DPI:     72,               // Точки на дюйм
		Hinting: font.HintingFull, // Качество рендеринга
	})

	return nil
}

// DrawText рисует текст в указанных координатах
func DrawText(screen *ebiten.Image, str string, x, y int, clr color.Color) {
	text.Draw(screen, str, regularFont, x, y, clr)
}

// DrawTextCentered рисует текст с центрированием
func DrawTextCentered(screen *ebiten.Image, str string, x, y int, clr color.Color) {
	bounds := text.BoundString(regularFont, str)
	x -= bounds.Max.X / 2
	y -= bounds.Max.Y / 2
	text.Draw(screen, str, regularFont, x, y, clr)
}
