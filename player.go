package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	speed         float64
	x             float64
	y             float64
	playerImage   *ebiten.Image
	Health        int
	Mana          int
	Exp           int
	inventory     *Inventory
	inventoryOpen bool
}

var inventoryToggle bool

func NewPlayer(screenWidth int, screenHeight int) *Player {
	playerImg := ebiten.NewImage(64, 64)                   // 64 x 64 Pixels
	playerImg.Fill(color.RGBA{R: 255, G: 0, B: 0, A: 255}) //color Red Test ONLY

	return &Player{
		inventory:     NewInventory(),
		inventoryOpen: false,
		speed:         1.5,
		x:             float64(screenWidth/2) - 16,
		y:             float64(screenHeight/2) - 16,
		playerImage:   playerImg,
		Health:        10,
		Mana:          10,
		Exp:           0,
	}
}

func (p *Player) UpdatePlayer() {
	//inventory stuff
	if ebiten.IsKeyPressed(ebiten.KeyI) && !inventoryToggle {
		p.inventoryOpen = !p.inventoryOpen
		inventoryToggle = true
	}

	if !ebiten.IsKeyPressed(ebiten.KeyI) {
		inventoryToggle = false
	}

	if p.inventoryOpen && ebiten.IsKeyPressed(ebiten.KeyEscape) {
		p.inventoryOpen = false
	}

	if p.inventoryOpen {
		p.inventory.UpdateInventory()
	}
}

func (p *Player) Draw(screen *ebiten.Image) {

	drawImageAt(screen, p.playerImage, p.x, p.y)

	if p.inventoryOpen {
		p.inventory.DrawInventory(screen)
	}
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
