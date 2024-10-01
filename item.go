package main

import "github.com/hajimehoshi/ebiten/v2"

type Item struct {
	itemImage *ebiten.Image
	itemCount int
	itemName  string
}

type Armor struct {
	Item
	Defense int
}

type Weapon struct {
	Item
	AttackPower int
}

type Consumable struct {
	Item
	Healing int
	Mana    int
}
