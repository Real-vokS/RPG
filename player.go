package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	speed       int
	x           int
	y           int
	playerImage *ebiten.Image
}

func NewPlayer(screenWidth int, screenHeight int) *Player {
	playerImg := ebiten.NewImage(32, 32)                   // 32 x 32 Pixels
	playerImg.Fill(color.RGBA{R: 255, G: 0, B: 0, A: 255}) //color Red Test ONLY
	return &Player{
		speed:       1,
		x:           (screenWidth / 2) - 16,
		y:           (screenWidth / 2) - 16,
		playerImage: playerImg,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(p.x), float64(p.y))
	screen.DrawImage(p.playerImage, opts)
}

func (p *Player) Move() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.y = p.y - p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.y = p.y + p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.x = p.x - p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.x = p.x + p.speed
	}

}
