package main

const (
	slotSize   = 64
	rows       = 4
	cols       = 6
	totalSlots = rows * col
)

type Inventory struct {
	slots     [totalSlots]*Item
	dragged   *Item
	dragIndex int
	mouseX    int
	mouseY    int
}

func NewInventory() Inventory {
	inv := Inventory{
		dragged:   nil,
		dragIndex: -1,
	}

	return inv
}

func (inv *Inventory) getSlotIndex(x, y int) int {
	if x < 0 || y < 0 || x >= cols*slotSize || y >= rows*slotSize {
		return -1
	}
	col := x / slotSize
	row := y / slotSize
	return row*cols + col

}
