package memento

type Item struct {
	id    int
	price int
}

func NewItem(id int, price int) *Item {
	return &Item{
		id:    id,
		price: price,
	}
}

func (i Item) Id() int {
	return i.id
}

func (i Item) Price() int {
	return i.price
}
