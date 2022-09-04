package composite

type Phage struct {
	BigItem
}

func NewCaulfieldsWarhammer(cost int) *Phage {
	return &Phage{
		BigItem: BigItem{
			name:              "CaulfieldsWarhammer",
			blacksmithingCost: cost,
		},
	}
}
