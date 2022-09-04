package composite

type RubyCrystal struct {
	SmallItem
}

func NewRubyCrystal(price int) *RubyCrystal {
	return &RubyCrystal{
		SmallItem: SmallItem{
			name:  "Ruby Crystal",
			price: price,
		},
	}
}
