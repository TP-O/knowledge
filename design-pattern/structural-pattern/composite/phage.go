package composite

type Phage struct {
	BigItem
}

func NewPhase(cost int) *Phage {
	return &Phage{
		BigItem: BigItem{
			name:              "Phage",
			blacksmithingCost: cost,
		},
	}
}
