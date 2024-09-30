package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

var player *Player

func (g *Game) Update() error {

	//Player Stuff
	player.Move()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	//player stuff
	player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {

	//player stuff
	player = NewPlayer()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Project A")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
