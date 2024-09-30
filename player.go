package main

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	speed int
	x     int
	y     int
}

func NewPlayer() *Player {
	return &Player{
		speed: 1,
		x:     200,
		y:     200,
	}
}

func (p *Player) PlayerMovement() {
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
