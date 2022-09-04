package composite

type LongSword struct {
	SmallItem
}

func NewLongSword(price int) *LongSword {
	return &LongSword{
		SmallItem: SmallItem{
			name:  "Long Sword",
			price: price,
		},
	}
}
