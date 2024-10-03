package main

type Enemy struct {
	speed float64
	x float64
	y float64
	enemyImage *ebiten.Image
	Health int
}

func newEnemy() *Enemy {
	enemyImg:= ebiten.NewImage(64,64) 				  // 64 x 64 pixels
	enemyImg.Fill(color{R: 0, G: 255, B:255, A: 255}) //only test

	return &Enemy{
		speed: 1.5,
		x: 200
		y: 200
		enemyImage: enemyImg,
		Health: 20,
	}
}

func (e *Enemy) DropItem {
	
}

func (e *Enemy) DrawEnemy {

}

func (e *Enemy) UpdateEnemy {

}