package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *Player
}

var screenWidth = 640
var screenHeight = 480

func NewGame() *Game {
	g := &Game{
		player: NewPlayer(screenWidth, screenHeight),
	}
	return g
}

func (g *Game) Update() error {

	//Player Stuff
	g.player.UpdatePlayer()
	g.player.Move()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	//player stuff
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Project A")

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
