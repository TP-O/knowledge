package composite

type Sheen struct {
	SmallItem
}

func NewSheen(price int) *Sheen {
	return &Sheen{
		SmallItem: SmallItem{
			name:  "Sheen",
			price: price,
		},
	}
}
