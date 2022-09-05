package memento

type InventoryState struct {
	balance int
	items   []Item
}

func NewInventoryState(blance int) *InventoryState {
	return &InventoryState{
		balance: blance,
	}
}

func (is InventoryState) Balance() int {
	return is.balance
}

func (is InventoryState) Items() []Item {
	return is.items
}
