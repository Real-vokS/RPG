package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	slotSize   = 64
	rows       = 4
	cols       = 6
	totalSlots = rows * cols
)

type Inventory struct {
	slots     [totalSlots]*Item
	dragged   *Item
	dragIndex int
	mouseX    int
	mouseY    int
}

func NewInventory() *Inventory {
	inv := &Inventory{
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

func (inv *Inventory) UpdateInventory() {
	inv.mouseX, inv.mouseY = ebiten.CursorPosition()

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if inv.dragged == nil {
			slotIndex := inv.getSlotIndex(inv.mouseX, inv.mouseY)
			if slotIndex >= 0 && inv.slots[slotIndex] != nil {
				inv.dragged = inv.slots[slotIndex]
				inv.dragIndex = slotIndex
				inv.slots[slotIndex] = nil
			}
		}
	} else if inv.dragged != nil {
		slotIndex := inv.getSlotIndex(inv.mouseX, inv.mouseY)
		if slotIndex >= 0 {
			if inv.slots[slotIndex] == nil {
				inv.slots[slotIndex] = inv.dragged
			} else if inv.slots[slotIndex].itemName == inv.dragged.itemName {
				inv.slots[slotIndex].itemCount += inv.dragged.itemCount
			} else {
				inv.slots[inv.dragIndex], inv.slots[slotIndex] = inv.slots[slotIndex], inv.dragged
			}
		} else {
			inv.slots[inv.dragIndex] = inv.dragged
		}
		inv.dragged = nil
		inv.dragIndex = -1
	}
}

func (inv *Inventory) DrawInventory(screen *ebiten.Image) {
	for i := 0; i < totalSlots; i++ {
		x := (i % cols) * slotSize
		y := (i / cols) * slotSize

		ebitenutil.DrawRect(screen, float64(x), float64(y), slotSize, slotSize, color.Gray{Y: 128})

		if inv.slots[i] != nil {
			drawImageAt(screen, inv.slots[i].itemImage, float64(x), float64(y))

			if inv.slots[i].itemCount > 1 {
				countStr := fmt.Sprintf("%d", inv.slots[1].itemCount)
				ebitenutil.DebugPrintAt(screen, countStr, x+5, y+5)
			}
		}
	}

	if inv.dragged != nil {
		drawImageAt(screen, inv.dragged.itemImage, float64(inv.mouseX-slotSize/2), float64(inv.mouseY-slotSize/2))

		if inv.dragged.itemCount > 1 {
			countStr := fmt.Sprintf("%d", inv.dragged.itemCount)
			ebitenutil.DebugPrintAt(screen, countStr, inv.mouseX-slotSize/2+5, inv.mouseY-slotSize/2+5)
		}
	}
}
