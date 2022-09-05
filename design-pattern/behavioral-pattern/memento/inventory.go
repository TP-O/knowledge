package memento

import (
	"fmt"
	"reflect"

	"golang.org/x/exp/slices"
)

type Inventory struct {
	*InventoryState

	history []InventoryState
}

func NewInventory(state *InventoryState) *Inventory {
	return &Inventory{
		InventoryState: state,
	}
}

func (i *Inventory) Add(item *Item) bool {
	fmt.Println("Buying item", item.Id(), "with", item.Price())

	if item.Price() > i.balance {
		fmt.Println("Don't enough money to buy item", item.Id())

		return false
	}

	i.Save()

	i.items = append(i.items, *item)
	i.balance -= item.Price()

	fmt.Println("Bought item", item.Id())

	return true
}

func (i *Inventory) Save() {
	if slices.IndexFunc(i.history, func(s InventoryState) bool {
		return reflect.DeepEqual(s, *i.InventoryState)
	}) == -1 {
		i.history = append(i.history, *i.InventoryState)
	}
}

func (i *Inventory) Undo() {
	if len(i.history) == 0 {
		return
	}

	prevState := i.history[len(i.history)-1]
	i.InventoryState = &prevState
	i.history = i.history[:len(i.history)-1]
}

func (i Inventory) Show() {
	fmt.Println("===")
	fmt.Println("Balance:", i.Balance())
	fmt.Println("Items:")

	for _, i := range i.Items() {
		fmt.Println("Item", i.Id())
	}

	fmt.Println("===")
}
