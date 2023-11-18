package inventoryUtil

import "github.com/kainn9/demo/components"

func AddItemToInventory(inventory *[]*components.InventoryItem, item *components.InventoryItem) {

	for _, inventoryItem := range *inventory {
		if inventoryItem.Name == item.Name {
			inventoryItem.Quantity += item.Quantity
			return // Item found and updated, no need to append
		}
	}

	// Item not found, append it to the inventory
	*inventory = append(*inventory, item)
}
func RemoveItemFromInventory(inventory *[]*components.InventoryItem, item *components.InventoryItem) {

	for i, inventoryItem := range *inventory {
		if inventoryItem.Name == item.Name {
			if inventoryItem.Quantity > item.Quantity {
				inventoryItem.Quantity -= item.Quantity

			} else if inventoryItem.Quantity == item.Quantity {
				*inventory = append((*inventory)[:i], (*inventory)[i+1:]...)

			}
		}
	}

}

func GetItemFromInventory(inventory []*components.InventoryItem, itemName string) *components.InventoryItem {

	for _, inventoryItem := range inventory {
		if inventoryItem.Name == itemName {
			return inventoryItem
		}
	}

	return nil
}
