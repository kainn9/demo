package components

import "github.com/yohamta/donburi"

var InventoryComponent = donburi.NewComponentType[[]*InventoryItem]()

type InventoryItem struct {
	Name     string
	Quantity int
}

func NewInventoryItem(name string, quantity int) *InventoryItem {
	return &InventoryItem{
		Name:     name,
		Quantity: quantity,
	}
}
