package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	inventory Inventory
}

var player *Player

func (g *Game) Update() error {

	//Player Stuff
	player.Move()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	//player stuff
	player.DrawPlayer(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {

	screenWidth := 640
	screenHeight := 480

	//player stuff
	player = NewPlayer(screenWidth, screenHeight)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Project A")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
