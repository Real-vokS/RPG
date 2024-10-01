package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func drawImageAt(screen *ebiten.Image, image *ebiten.Image, x float64, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(image, op)
}
